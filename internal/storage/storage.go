package storage

import "errors"

var (
	ErrPCNotFound          = errors.New("pc not found")
	ErrPCAlreadyExists     = errors.New("pc already exists")
	ErrRamNotFound         = errors.New("ram not found")
	ErrRamAlreadyExists    = errors.New("ram already exists")
	ErrCpuAlreadyExists    = errors.New("cpu already exists")
	ErrCpuNotFound         = errors.New("cpu not found")
	ErrGpuAlreadyExists    = errors.New("gpu already exists")
	ErrGpuNotFound         = errors.New("gpu not found")
	ErrMemoryAlreadyExists = errors.New("memory already exists")
	ErrMemoryNotFound      = errors.New("memory not found")
)
