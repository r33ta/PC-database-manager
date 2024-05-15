package main

import (
	"log/slog"
	"os"

	"github.com/r33ta/pc-database-manager/internal/config"
	"github.com/r33ta/pc-database-manager/internal/lib/logger/sl"
	"github.com/r33ta/pc-database-manager/internal/models/memory"
	"github.com/r33ta/pc-database-manager/internal/models/pc_storage"
	"github.com/r33ta/pc-database-manager/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting...", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	// testing
	memoryID, err := storage.SaveMemory("Kingston RAM", memory.DDR4, 32)
	if err != nil {
		log.Error("failed to save memory", sl.Err(err))
	}
	cpuID, err := storage.SaveCpu("Intel Core i7-9700K", 8, 8, 3600)
	if err != nil {
		log.Error("failed to save cpu", sl.Err(err))
	}
	gpuID, err := storage.SaveGpu("Nvidia RTX 2080", "ASUS", 8, 1800)
	if err != nil {
		log.Error("failed to save gpu", sl.Err(err))
	}
	storageID, err := storage.SaveStorage("Kingston SSD 1TB", 1000, pc_storage.SSD)
	if err != nil {
		log.Error("failed to save storage", sl.Err(err))
	}
	_, err = storage.SavePC("Some PC", memoryID, cpuID, gpuID, storageID)
	if err != nil {
		log.Error("failed to save pc", sl.Err(err))
	}

	// TODO: init router: chi, "chi render"

	// TODO: start server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
