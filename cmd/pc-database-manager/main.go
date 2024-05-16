package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/r33ta/pc-database-manager/internal/config"
	"github.com/r33ta/pc-database-manager/internal/lib/logger/sl"
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
	pcs, err := storage.GetPC(1)
	if err != nil {
		log.Error("failed to get all pcs", sl.Err(err))
		os.Exit(1)
	}
	fmt.Println(pcs)

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
