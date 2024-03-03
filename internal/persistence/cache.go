package persistence

import (
	"database/sql"
	"fmt"
)

type Cache struct {
	db *sql.DB
}

func New(path string) (*Cache, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("Can't open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Can't connect to database: %w", err)
	}
	return &Cache{db: db}, nil
}
