package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ismailozdel/micro/invoice/handler"
	"github.com/ismailozdel/micro/invoice/services"
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

	stockGrpcConn := NewGrpcClient(os.Getenv("STOCK_HOST"))
	defer stockGrpcConn.Close()

	userGrpcConn := NewGrpcClient(os.Getenv("USER_HOST"))
	defer userGrpcConn.Close()

	invoiceService := services.NewInvoiceService(stockGrpcConn, userGrpcConn)
	handler.NewInvoiceGrpcHandler(grpcServer, invoiceService)

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

func NewGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
