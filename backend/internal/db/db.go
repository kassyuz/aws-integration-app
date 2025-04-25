// File: backend/internal/db/db.go
package db

import (
	"database/sql"
	"fmt"

	"github.com/yourname/aws-integration-app/internal/config"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Connect establishes a connection to the database
func Connect(cfg config.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// File: backend/internal/db/migrations.go
package db

import (
	"database/sql"
	"fmt"
)

// RunMigrations executes all database migrations
func RunMigrations(db *sql.DB) error {
	// Create users table
	if err := createUsersTable(db); err != nil {
		return err
	}

	// Add more migrations as needed

	return nil
}

// createUsersTable creates the users table if it doesn't exist
func createUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		company TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		phone TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	return nil
}