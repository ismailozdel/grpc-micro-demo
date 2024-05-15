package main

type Config struct {
	Host string `envconfig:"HOST" default:"127.0.0.1"`
	Port string `envconfig:"PORT" default:"5000"`
}
