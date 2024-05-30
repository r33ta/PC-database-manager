package main

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/r33ta/pc-database-manager/internal/config"
	"github.com/r33ta/pc-database-manager/internal/http-server/handlers/cpu/savecpu"
	"github.com/r33ta/pc-database-manager/internal/http-server/handlers/gpu/savegpu"
	"github.com/r33ta/pc-database-manager/internal/http-server/handlers/memory/savememory"
	"github.com/r33ta/pc-database-manager/internal/http-server/handlers/pc/savepc"
	"github.com/r33ta/pc-database-manager/internal/http-server/handlers/ram/saveram"
	mwLogger "github.com/r33ta/pc-database-manager/internal/http-server/middleware/logger"
	"github.com/r33ta/pc-database-manager/internal/lib/logger/handlers/slogpretty"
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
	log.Error("error messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/save/pc", savepc.NewPC(log, storage))
	router.Post("/save/ram", saveram.NewRAM(log, storage))
	router.Post("/save/cpu", savecpu.NewCPU(log, storage))
	router.Post("/save/gpu", savegpu.NewGPU(log, storage))
	router.Post("/save/memory", savememory.NewMemory(log, storage))

	log.Info("starting server", slog.String("address", cfg.Address))

	// TODO: start server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
