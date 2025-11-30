// Package main contains the backend server logic for TaskBoard, including
// database initialization, environment variable handling, and application startup.
package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global GORM database connection used throughout the backend.
var DB *gorm.DB

// initDB initializes the PostgreSQL connection using environment variables
// and runs automatic migrations for all database models.
func initDB() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "taskboard")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db
	log.Println("✅ Connected to PostgreSQL and migrated")
}

// getEnv returns an environment variable value or a default
// value if the variable is not set.
func getEnv(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return def
	}
	return val
}
