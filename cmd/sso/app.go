package sso

import (
	"log/slog"
	"os"
	"sso/internal/config"
	"sso/internal/domain"
)

func Run() {
	// config
	cfg := config.MustLoad()
	// logger
	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.String("env", cfg.Env), slog.Any("config", cfg))

	// todo app
	// todo run grpc
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	// привязываем уровень логирования к окружению
	switch env {
	case domain.EnvLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case domain.EnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case domain.EnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
