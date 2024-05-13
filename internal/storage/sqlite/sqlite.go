package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
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
		СREATE TABLE IF NOT EXISTS memory (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
			memory_type TEXT NOT NULL
			capacity INTEGER NOT NULL
		)
		CREATE TABLE IF NOT EXISTS cpu (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
			cores INTEGER NOT NULL
			threads INTEGER NOT NULL
			frequency FLOAT NOT NULL
		)
		CREATE TABLE IF NOT EXISTS gpu (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
			vendor TEXT NOT NULL
			memory INTEGER NOT NULL
			frequency FLOAT NOT NULL
		)
		CREATE TABLE IF NOT EXISTS storage (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
			type TEXT NOT NULL
			capacity INTEGER NOT NULL
			speed INTEGER NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if _, err := stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
