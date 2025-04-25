// File: backend/internal/config/config.go
package config

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port         string
	AllowOrigins []string
}

// DatabaseConfig holds database connection details
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig holds JWT token configuration
type JWTConfig struct {
	Secret    string
	ExpiresIn time.Duration
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()

	// Load JWT configuration
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET environment variable is required")
	}

	// Default to 24 hours if not specified
	jwtExpiresIn := 24 * time.Hour

	// Database configuration - default to environment variables
	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "mcpapp"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Server configuration
	serverConfig := ServerConfig{
		Port: getEnv("PORT", "8080"),
		AllowOrigins: []string{
			getEnv("ALLOW_ORIGIN", "http://localhost:3000"),
		},
	}

	return &Config{
		Server:   serverConfig,
		Database: dbConfig,
		JWT: JWTConfig{
			Secret:    jwtSecret,
			ExpiresIn: jwtExpiresIn,
		},
	}, nil
}

// Helper function to get an environment variable or return a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}