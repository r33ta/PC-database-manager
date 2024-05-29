package storage

import "errors"

var (
	ErrPCNotFound          = errors.New("pc not found")
	ErrPCAlreadyExists     = errors.New("pc already exists")
	ErrRAMNotFound         = errors.New("ram not found")
	ErrRAMAlreadyExists    = errors.New("ram already exists")
	ErrCPUAlreadyExists    = errors.New("cpu already exists")
	ErrCPUNotFound         = errors.New("cpu not found")
	ErrGPUAlreadyExists    = errors.New("gpu already exists")
	ErrGPUNotFound         = errors.New("gpu not found")
	ErrMemoryAlreadyExists = errors.New("memory already exists")
	ErrMemoryNotFound      = errors.New("memory not found")
)
