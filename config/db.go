package config

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/pressly/goose/v3"
)

// Run latest migrations
func RunMigrations(db *sql.DB) {
	migrationsDir, err := filepath.Abs("../migrations")
	if err != nil {
		log.Fatal("Error resolving migrations path: ", err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatal("Error Running Migrations", err)
	}
}
