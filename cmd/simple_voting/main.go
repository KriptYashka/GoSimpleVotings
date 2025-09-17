package main

import (
	"log/slog"
	"os"
	"simple_voting/internal/config"
	"simple_voting/internal/lib/logger/handlers/slogdiscard/sl"
	"simple_voting/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env)) // к каждому сообщению будет добавляться поле с информацией о текущем окружении

	log.Info("initializing server", slog.String("address", cfg.Address)) // Помимо сообщения выведем параметр с адресом
	log.Debug("logger debug mode enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	var level slog.Level

	switch env {
	case envLocal:
		level = slog.LevelDebug
	case envDev:
		level = slog.LevelDebug
	case envProd:
		level = slog.LevelInfo
	}
	log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	return log
}
