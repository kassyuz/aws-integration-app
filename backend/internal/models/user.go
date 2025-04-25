// File: backend/internal/models/user.go
package models

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Company   string    `json:"company"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"-"` // Never sent to client
	CreatedAt time.Time `json:"created_at"`
}

// UserRegistration represents the registration payload
type UserRegistration struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// CreateUser creates a new user in the database
func CreateUser(db *sql.DB, reg UserRegistration) (int, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("failed to hash password: %w", err)
	}

	// Insert into database
	var id int
	err = db.QueryRow(`
		INSERT INTO users (name, company, email, phone, password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, reg.Name, reg.Company, reg.Email, reg.Phone, string(hashedPassword)).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var user User
	var hashedPassword string

	err := db.QueryRow(`
		SELECT id, name, company, email, phone, password, created_at
		FROM users WHERE email = $1
	`, email).Scan(&user.ID, &user.Name, &user.Company, &user.Email, &user.Phone, &hashedPassword, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	user.Password = hashedPassword
	return &user, nil
}

// ValidatePassword checks if the provided password matches the user's hashed password
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}