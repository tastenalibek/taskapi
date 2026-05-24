package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(dsn string) (*sql.DB, error) {
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}
	database.SetMaxOpenConns(25)
	database.SetMaxIdleConns(5)
	return database, nil
}

func Migrate(database *sql.DB) error {
	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id          SERIAL PRIMARY KEY,
			title       TEXT NOT NULL,
			description TEXT NOT NULL DEFAULT '',
			status      TEXT NOT NULL DEFAULT 'todo'
				CHECK (status IN ('todo', 'in_progress', 'done')),
			created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
			updated_at  TIMESTAMPTZ NOT NULL DEFAULT now()
		)
	`)
	return err
}
