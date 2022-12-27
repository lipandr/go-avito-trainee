package internal

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:5000"`
}

func InitConfig() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&cfg.ServerAddress, "s", cfg.ServerAddress, "Server address")
	flag.Parse()

	return cfg
}
