package config

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

// Run latest migrations
func RunMigrations(db *sql.DB) {
	if err := goose.Up(db, "../migrations"); err != nil {
		log.Fatal("Error Running Migrations")
	}
}
