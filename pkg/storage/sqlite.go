package storage

import (
	"database/sql"
	"fmt"

	"github.com/xSaCh/dweep/pkg/storage/sqlite"
)

type SqliteStore struct {
	db *sql.DB
	sqlite.SqliteWLStore
}

func NewSqliteStore(file string) (*SqliteStore, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %v", err)
	}
	s := &SqliteStore{db: db, SqliteWLStore: *sqlite.NewSqlWLStore(db)}

	return s, nil
}

func (s *SqliteStore) Create() error {
	err := s.WLCreate()
	if err != nil {
		return fmt.Errorf("error creating watchlist table: %v", err)
	}
	return nil
}
