package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) error {
	// Get the absolute path to the migrations directory
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("could not get current file path")
	}
	migrationsPath := filepath.Join(filepath.Dir(filename))

	// Convert Windows path to forward slashes for the file:// URL
	migrationsURL := "file://" + filepath.ToSlash(migrationsPath)

	// Create database driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("could not create migration driver: %w", err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("could not create migration instance: %w", err)
	}

	// Run migrations
	log.Println("Running database migrations...")
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}
