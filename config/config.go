package config

import (
	"log/slog"
	"sync"

	env "github.com/caarlos0/env/v10"
)

type config struct {
	Port string `env:"PORT"`
}

var (
	Val  config
	once sync.Once
)

func init() {
	once.Do(
		func() {
			cfg := config{}
			if err := env.Parse(&cfg); err != nil {
				slog.Error(err.Error())
			}

			Val = cfg
		},
	)
}
