package postgres

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(dbstring string) (*Storage, error) {
	db, err := sqlx.Connect("postgres", dbstring)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to postgres '%s'", dbstring)
	}
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	return &Storage{db: db}, nil
}

func(s *Storage) Migrate(dir string) error {
	return goose.Run("up", s.db.DB, dir)
}

func NewTestStorage(dbstring string, migrationDir string) (*Storage, func()) {
	db, teardown := MustNewDevelopmentDB(dbstring, migrationDir)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(time.Hour)

	return &Storage{db: db}, teardown
}