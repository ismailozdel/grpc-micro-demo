package main

import (
	"log"

	"google.golang.org/grpc"
)

func NewGrpcClient(addr string) *grpc.ClientConn {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
