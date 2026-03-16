package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"golang-api/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to database
	db, err := connectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run seeds
	if err := runSeeds(db); err != nil {
		log.Fatalf("Failed to run seeds: %v", err)
	}

	log.Println("✅ Seeds completed successfully")
}

func connectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)
	return sql.Open("pgx", dsn)
}

func runSeeds(db *sql.DB) error {
	// Read seed file
	seedSQL, err := os.ReadFile("migrations/seeds/001_seed_users.sql")
	if err != nil {
		return fmt.Errorf("failed to read seed file: %w", err)
	}

	// Execute seed SQL
	_, err = db.Exec(string(seedSQL))
	if err != nil {
		return fmt.Errorf("failed to execute seeds: %w", err)
	}

	return nil
}
