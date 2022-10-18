package database

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed migration/schema.up.sql
var schemaUp string

//go:embed migration/schema.down.sql
var schemaDown string

// IDatabase - The interface for the database driver
type IDatabase interface {
	Ping() error
	Migrate(cmd string) error
	Close() error
}

// Database - The database driver struct
type Database struct {
	DB *sql.DB
}

// NewDatabase - Creates a new connection to the database
func NewDatabase(dbUser, dbPassword, dbHost, dbPort, dbName, dbDriver string) (*Database, error) {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open(dbDriver, DBURL)
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

// Migrate - Creates the tables if we pass `up` as the argument.
// Removes the tables if we pass `down` as the argument.
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

// Ping - Checks database health
func (s *Database) Ping() error {
	return s.DB.Ping()
}

// Close - Closes the connection to the database
func (s *Database) Close() error {
	return s.DB.Close()
}
