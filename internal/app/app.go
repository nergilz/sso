package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
)

type AppUsecase struct {
	GrpcSrv *grpcapp.App
}

func New(l *slog.Logger, p int, storagePath, tokenTTL string) *AppUsecase {
	// todo: init storage
	// todo: init auth service

	grpcApp := grpcapp.New(l, p)

	return &AppUsecase{
		GrpcSrv: grpcApp,
	}
}
