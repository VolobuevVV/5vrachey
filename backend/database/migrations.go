package database

import (
	"errors"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Fatal("Error creating migration driver:", err)
	}

	migrationsPath := "file://database/migrations"
	if _, err := os.Stat("/app/backend/database/migrations"); err == nil {
		migrationsPath = "file:///app/backend/database/migrations"
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		log.Fatal("Error creating migration instance:", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Error applying migrations:", err)
	}

	log.Println("Migrations applied successfully!")
}
