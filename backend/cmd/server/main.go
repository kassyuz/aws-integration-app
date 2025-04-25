// File: backend/cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yourname/aws-integration-app/internal/api"
	"github.com/yourname/aws-integration-app/internal/config"
	"github.com/yourname/aws-integration-app/internal/db"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	database, err := db.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Run migrations
	if err := db.RunMigrations(database); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	// Set up API server
	router := api.SetupRoutes(database, cfg)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
