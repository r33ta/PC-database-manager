package storage

import "errors"

var (
	ErrPCNotFound           = errors.New("pc not found")
	ErrPCAlreadyExists      = errors.New("pc already exists")
	ErrMemoryNotFound       = errors.New("memory not found")
	ErrCpuAlreadyExists     = errors.New("cpu already exists")
	ErrCpuNotFound          = errors.New("cpu not found")
	ErrGpuAlreadyExists     = errors.New("gpu already exists")
	ErrGpuNotFound          = errors.New("gpu not found")
	ErrStorageAlreadyExists = errors.New("storage already exists")
	ErrStorageNotFound      = errors.New("storage not found")
	ErrMemoryAlreadyExists  = errors.New("memory already exists")
)
