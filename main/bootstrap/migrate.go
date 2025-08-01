package bootstrap

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sqlx.DB) {
	migrationDir := "./migrations"

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	if err := goose.Up(db.DB, migrationDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
