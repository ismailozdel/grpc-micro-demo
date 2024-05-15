package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	config := Config{}
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := NewGRPCServer(":" + config.GRPCPort)
	gRPCServer.Run()
}
