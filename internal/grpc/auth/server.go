package auth

import (
	"context"
	"fmt"
	"log"

	ssov1 "github.com/nergilz/protos-url-shortiner/gen/go/sso"
	"google.golang.org/grpc"
)

/*
   реализуем интерфейс grpc сервера
*/

type serverAPI struct {
	// для запуска сервера без реализации методов интерфейса
	ssov1.UnimplementedAuthServer
}

func Register(grpcServer *grpc.Server) {
	ssov1.RegisterAuthServer(grpcServer, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	// log.Fatalf("test login method")
	fmt.Println("login by grpc")

	return &ssov1.LoginResponse{
		Token: fmt.Sprintf("tets.token.with.email.%s", req.GetEmail()),
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	log.Fatalf("test register method")

	return nil, nil
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	log.Fatalf("test isadmin method")

	return nil, nil
}
