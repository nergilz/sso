package sso

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/domain"
	"sso/internal/lib/logger/slogpretty"
	"syscall"
)

func Run() {
	// config
	cfg := config.MustLoad()
	// logger
	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.String("env", cfg.Env), slog.Any("config", cfg))

	// app
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL.String())

	go application.GrpcSrv.MustRun()

	// shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	// чтение из канала блокирующая операция, ждем сигнал от ОС
	sig := <-stop
	application.GrpcSrv.Stop(sig)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	// привязываем уровень логирования к окружению
	switch env {
	case domain.EnvLocal:
		log = setupPrettySlog()
		// log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case domain.EnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case domain.EnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
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
