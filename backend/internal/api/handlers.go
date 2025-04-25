// File: backend/internal/api/handlers.go
package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/yourname/aws-integration-app/internal/auth"
	awsClient "github.com/yourname/aws-integration-app/internal/aws"
	"github.com/yourname/aws-integration-app/internal/config"
	"github.com/yourname/aws-integration-app/internal/models"
)

// registerHandler handles user registration
func registerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reg models.UserRegistration
		if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Create the user
		_, err := models.CreateUser(db, reg)
		if err != nil {
			// Check for duplicate email (this check is basic, improve in production)
			if err.Error() == "failed to create user: pq: duplicate key value violates unique constraint \"users_email_key\"" {
				respondWithError(w, http.StatusConflict, "Email already registered")
				return
			}
			respondWithError(w, http.StatusInternalServerError, "Failed to create user")
			return
		}

		respondWithJSON(w, http.StatusCreated, map[string]string{
			"message": "User registered successfully",
		})
	}
}

// loginHandler handles user login
func loginHandler(db *sql.DB, cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Get user by email
		user, err := models.GetUserByEmail(db, creds.Email)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		// Validate password
		if !user.ValidatePassword(creds.Password) {
			respondWithError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		// Generate JWT token
		token, err := auth.GenerateToken(user, cfg.JWT)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to generate token")
			return
		}

		// Return token and user info
		respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"token": token,
			"user":  user,
		})
	}
}

// verifyAwsHandler verifies AWS credentials
func verifyAwsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var awsCreds awsClient.Credentials
		if err := json.NewDecoder(r.Body).Decode(&awsCreds); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Verify the AWS credentials
		if err := awsClient.VerifyCredentials(awsCreds); err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Here you could store a reference to these credentials in your database
		// but NEVER store the actual secret key

		respondWithJSON(w, http.StatusOK, map[string]string{
			"message": "AWS credentials verified successfully",
		})
	}
}

// Helper function to respond with an error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}