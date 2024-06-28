package grpcapp

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	authGRPC "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	Log        *slog.Logger
	GrpcServer *grpc.Server
	Port       int
}

func New(l *slog.Logger, p int) *App {
	server := grpc.NewServer()

	// подключение обработчика
	authGRPC.Register(server)

	return &App{
		Log:        l,
		Port:       p,
		GrpcServer: server,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		a.Log.Warn("can not run grpc", slog.String("error", err.Error()))
		log.Fatalf("panic by grpc server")
	}
}

func (a *App) Run() error {
	const op = "grpcapp.run" // сокращение от operation
	log := a.Log.With(slog.String("operation", op), slog.Int("port", a.Port))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		return fmt.Errorf("%s, %s", op, err.Error())
	}

	log.Info("grpc app server is running") //, slog.String("addr", listener.Addr().String()))

	if err := a.GrpcServer.Serve(listener); err != nil {
		return fmt.Errorf("%s, %s", op, err.Error())
	}

	return nil
}

func (a *App) Stop(sig os.Signal) {
	const op = "grpcapp.stop"

	a.Log.With(slog.String("operation", op), slog.Any("os signal", sig), slog.Int("port", a.Port)).Info("graceful stopped grpc server", slog.Int("port", a.Port))

	a.GrpcServer.GracefulStop()
}
