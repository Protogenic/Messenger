package app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"messenger/internal/auth"
	"net"
)

type App struct {
	port       int
	grpcServer *grpc.Server
}

func (a *App) Run() {
	client, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if a.grpcServer.Serve(client) != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *App) Stop() {
	a.grpcServer.GracefulStop()
}

func New(port int) *App {
	grpcServer := grpc.NewServer()
	auth.Register(grpcServer)
	return &App{
		port:       port,
		grpcServer: grpcServer,
	}

}
