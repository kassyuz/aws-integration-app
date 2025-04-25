// File: backend/internal/api/middleware.go
package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/yourname/aws-integration-app/internal/auth"
	"github.com/yourname/aws-integration-app/internal/config"
)

// Key type for context values
type contextKey string

// Context keys
const (
	userIDKey contextKey = "userID"
	emailKey  contextKey = "email"
)

// AuthMiddleware validates JWT tokens and sets user information in request context
func AuthMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				respondWithError(w, http.StatusUnauthorized, "Authorization header required")
				return
			}

			// Extract the token from the Authorization header
			// The header format should be "Bearer {token}"
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				respondWithError(w, http.StatusUnauthorized, "Invalid authorization format")
				return
			}

			tokenString := tokenParts[1]

			// Validate the token
			claims, err := auth.ValidateToken(tokenString, cfg.JWT)
			if err != nil {
				respondWithError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			// Add claims to request context
			ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
			ctx = context.WithValue(ctx, emailKey, claims.Email)

			// Call the next handler with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserIDFromContext extracts the user ID from the request context
func GetUserIDFromContext(r *http.Request) (int, bool) {
	id, ok := r.Context().Value(userIDKey).(int)
	return id, ok
}

// GetEmailFromContext extracts the email from the request context
func GetEmailFromContext(r *http.Request) (string, bool) {
	email, ok := r.Context().Value(emailKey).(string)
	return email, ok
}