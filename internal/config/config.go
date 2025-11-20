package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Cheemx/go-hrms/internal/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

type APIConfig struct {
	DB        *database.Queries
	Scheduler *cron.Cron
}

func Load() *APIConfig {
	// Load Environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: .env file not found!")
	}

	cron := cron.New()

	// Initialize DB
	dbURL := mustGetEnv("DB_URL")
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	dbQueries := database.New(db)
	cfg := &APIConfig{
		DB:        dbQueries,
		Scheduler: cron,
	}

	fmt.Println("MySQL Database Connected Successfully.")
	return cfg
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("environment variable %s not set", key)
	}
	return val
}
