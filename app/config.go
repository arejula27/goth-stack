package main

import (
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type config struct {
	Addres       string `env:"HTTP_LISTEN_ADDR" envDefault:":3000"`
	IsProduction bool   `env:"PRODUCTION"`
}

func loadConfig() config {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal(err)
	}
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}
	return cfg

}
