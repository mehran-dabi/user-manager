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

type IDatabase interface {
	Ping() error
	Migrate(cmd string) error
	Close() error
}

type Database struct {
	DB *sql.DB
}

// NewDatabase example: ./tmp/
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
