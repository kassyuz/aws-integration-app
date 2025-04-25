// File: backend/internal/api/routes.go
package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/yourname/aws-integration-app/internal/config"
)

// SetupRoutes configures all the API routes
func SetupRoutes(db *sql.DB, cfg *config.Config) http.Handler {
	router := mux.NewRouter()

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Public routes - no authentication required
	api.HandleFunc("/register", registerHandler(db)).Methods("POST")
	api.HandleFunc("/login", loginHandler(db, cfg)).Methods("POST")

	// Protected routes - require authentication
	protected := api.PathPrefix("").Subrouter()
	protected.Use(AuthMiddleware(cfg))
	protected.HandleFunc("/verify-aws", verifyAwsHandler(db)).Methods("POST")

	// Add more routes as needed

	// CORS handling
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins(cfg.Server.AllowOrigins),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return corsHandler(router)
}