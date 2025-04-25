// File: backend/internal/auth/jwt.go
package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yourname/aws-integration-app/internal/config"
	"github.com/yourname/aws-integration-app/internal/models"
)

// Claims represents the JWT claims
type Claims struct {
	UserID int    `json:"id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(user *models.User, cfg config.JWTConfig) (string, error) {
	expirationTime := time.Now().Add(cfg.ExpiresIn)

	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.Secret))

	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken verifies that a token is valid and returns the claims
func ValidateToken(tokenString string, cfg config.JWTConfig) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.Secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}