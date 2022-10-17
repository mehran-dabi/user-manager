package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed migration/schema.up.sql
var schemaUp string

//go:embed migration/schema.down.sql
var schemaDown string

type IDatabase interface {
	Ping() error
	Migrate(cmd string) error
	Close() error
}

type Database struct {
	DB *sql.DB
}

// NewDatabase example: ./tmp/
func NewDatabase(dir string) (*Database, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	storeFile := path.Join(dir, "store.db")

	// create file if not exist
	if _, err := os.Stat(storeFile); os.IsNotExist(err) {
		f, err := os.Create(storeFile)
		if err != nil {
			return nil, err
		}
		_ = f.Close()
	}

	db, err := sql.Open("sqlite3", storeFile+"?cache=shared_sync=1&_cache_size=25000")
	if err != nil {
		return nil, err
	}

	// check connectivity
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}

// Migrate does the migration of the tables
func (s *Database) Migrate(cmd string) error {
	switch cmd {
	case "up":
		_, err := s.DB.ExecContext(context.Background(), schemaUp)
		return err
	case "down":
		_, err := s.DB.ExecContext(context.Background(), schemaDown)
		return err
	default:
		return fmt.Errorf("unknown command")
	}
}

// Ping check database ping
func (s *Database) Ping() error {
	return s.DB.Ping()
}

// Close closes the connection to the database
func (s *Database) Close() error {
	return s.DB.Close()
}
