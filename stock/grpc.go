package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ismailOZdel/micro/stock/handlers"
	"github.com/ismailOZdel/micro/stock/services"
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
	stockService := services.NewStockService()
	handlers.NewGrpcStockService(grpcServer, stockService)

	log.Printf("stock gRPC server listening on %s", s.addr)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		<-ctx.Done()
		grpcServer.GracefulStop()
		lis.Close()
		log.Println("stock gRPC server stopped")
	}()

	return grpcServer.Serve(lis)
}
