package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ismailozdel/micro/user/handlers"
	"github.com/ismailozdel/micro/user/services"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	userService := services.NewUserService()
	handlers.NewGrpcUsersService(grpcServer, userService)

	log.Printf("user gRPC server listening on %s", s.addr)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
		lis.Close()
		log.Println("user gRPC server stopped")
	}()

	return grpcServer.Serve(lis)
}
