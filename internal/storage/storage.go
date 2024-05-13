package storage

import "errors"

var (
	ErrPCNotFound      = errors.New("pc not found")
	ErrMemoryNotFound  = errors.New("memory not found")
	ErrCpuNotFound     = errors.New("cpu not found")
	ErrGpuNotFound     = errors.New("gpu not found")
	ErrStorageNotFound = errors.New("storage not found")
)
