package storage

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(StoragePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", StoragePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS pc (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			memory_id INTEGER NOT NULL,
			cpu_id INTEGER NOT NULL,
			gpu_id INTEGER NOT NULL,
			storage_id INTEGER NOT NULL)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
}
