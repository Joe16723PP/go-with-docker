package config

import (
	"github.com/joeshaw/envdecode"
	"log"
	"time"
)

type Conf struct {
	Server serverConf
	Debug  bool `env:"DEBUG,required"`
}

// property in .env file
type serverConf struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.Decode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
