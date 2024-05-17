package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"github.com/r33ta/pc-database-manager/internal/models/cpu"
	"github.com/r33ta/pc-database-manager/internal/models/gpu"
	"github.com/r33ta/pc-database-manager/internal/models/memory"
	"github.com/r33ta/pc-database-manager/internal/models/pc"
	"github.com/r33ta/pc-database-manager/internal/models/ram"
	"github.com/r33ta/pc-database-manager/internal/storage"
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
			ram_id INTEGER NOT NULL,
			cpu_id INTEGER NOT NULL,
			gpu_id INTEGER NOT NULL,
			storage_id INTEGER NOT NULL,
			FOREIGN KEY(ram_id) REFERENCES ram(id),
			FOREIGN KEY(cpu_id) REFERENCES cpu(id),
			FOREIGN KEY(gpu_id) REFERENCES gpu(id),
			FOREIGN KEY(storage_id) REFERENCES storage(id)
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if _, err := stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = db.Prepare(
		`CREATE TABLE IF NOT EXISTS ram (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			memory_type TEXT NOT NULL,
			capacity INTEGER NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if _, err := stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = db.Prepare(
		`CREATE TABLE IF NOT EXISTS cpu (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			cores INTEGER NOT NULL,
			threads INTEGER NOT NULL,
			frequency INTEGER NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if _, err := stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = db.Prepare(
		`CREATE TABLE IF NOT EXISTS gpu (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			manufacturer TEXT NOT NULL,
			memory INTEGER NOT NULL,
			frequency INTEGER NOT NULL
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if _, err := stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err = db.Prepare(
		`CREATE TABLE IF NOT EXISTS storage (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			capacity INTEGER NOT NULL,
			type TEXT NOT NULL
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

func (s *Storage) SavePC(name string, ramID, cpuID, gpuID, storageID int64) (int64, error) {
	const op = "storage.sqlite.SavePC"

	stmt, err := s.db.Prepare("INSERT INTO pc (name, ram_id, cpu_id, gpu_id, storage_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.Exec(name, ramID, cpuID, gpuID, storageID)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrPCAlreadyExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) SaveRAM(name string, memoryType string, capacity int64) (int64, error) {
	const op = "storage.sqlite.SaveRam"

	stmt, err := s.db.Prepare("INSERT INTO ram (name, memory_type, capacity) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.Exec(name, memoryType, capacity)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrRamAlreadyExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) SaveCpu(name string, cores, threads, frequency int64) (int64, error) {
	const op = "storage.sqlite.SaveCpu"

	stmt, err := s.db.Prepare("INSERT INTO cpu (name, cores, threads, frequency) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.Exec(name, cores, threads, frequency)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrCpuAlreadyExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) SaveGpu(name string, manufacturer string, memory, frequency int64) (int64, error) {
	const op = "storage.sqlite.SaveGpu"

	stmt, err := s.db.Prepare("INSERT INTO gpu (name, manufacturer, memory, frequency) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.Exec(name, manufacturer, memory, frequency)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrGpuAlreadyExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) SaveMemory(name string, capacity int64, storage_type string) (int64, error) {
	const op = "storage.sqlite.SaveMemory"

	stmt, err := s.db.Prepare("INSERT INTO storage (name, capacity, type) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	res, err := stmt.Exec(name, capacity, storage_type)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrMemoryAlreadyExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetPC(id int64) (*pc.PC, error) {
	const op = "storage.sqlite.GetPC"

	stmt, err := s.db.Prepare("SELECT name, ram_id, cpu_id, gpu_id, storage_id FROM pc WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var name string
	var ramID, cpuID, gpuID, storageID int64
	err = stmt.QueryRow(id).Scan(&name, &ramID, &cpuID, &gpuID, &storageID)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return &pc.PC{ID: id, Name: name, RAMID: ramID, CPUID: cpuID, GPUID: gpuID, StorageID: storageID}, nil
}

func (s *Storage) GetCpu(id int64) (*cpu.CPU, error) {
	const op = "storage.sqlite.GetCpu"

	stmt, err := s.db.Prepare("SELECT name, cores, threads, frequency FROM cpu WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var name string
	var cores, threads, frequency int64
	err = stmt.QueryRow(id).Scan(&name, &cores, &threads, &frequency)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return &cpu.CPU{ID: id, Name: name, Cores: cores, Threads: threads, Frequency: frequency}, nil
}

func (s *Storage) GetGpu(id int64) (*gpu.GPU, error) {
	const op = "storage.sqlite.GetGpu"

	stmt, err := s.db.Prepare("SELECT name, manufacturer, memory, frequency FROM gpu WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var name, manufacturer string
	var memory, frequency int64
	err = stmt.QueryRow(id).Scan(&name, &manufacturer, &memory, &frequency)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return &gpu.GPU{ID: id, Name: name, Manufacturer: manufacturer, Memory: memory, Frequency: frequency}, nil
}

func (s *Storage) GetRam(id int64) (*ram.RAM, error) {
	const op = "storage.sqlite.GetRam"

	stmt, err := s.db.Prepare("SELECT name, memory_type, capacity FROM memory WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var name, memoryType string
	var capacity int64
	err = stmt.QueryRow(id).Scan(&name, &memoryType, &capacity)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return &ram.RAM{ID: id, Name: name, MemoryType: memoryType, Capacity: capacity}, nil
}

func (s *Storage) GetMemory(id int64) (*memory.Memory, error) {
	const op = "storage.sqlite.GetMemory"

	stmt, err := s.db.Prepare("SELECT name, capacity, type FROM storage WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var name, storageType string
	var capacity int64
	err = stmt.QueryRow(id).Scan(&name, &capacity, &storageType)
	if err != nil {
		return nil, fmt.Errorf("%s: execute statement: %w", op, err)
	}

	return &memory.Memory{ID: id, Name: name, Capacity: capacity, StorageType: storageType}, nil
}
