package config

import (
	"log/slog"
	"sync"

	env "github.com/caarlos0/env/v10"
)

type config struct {
	Port          string `env:"PORT"`
	DBHost        string `env:"DBHOST"`
	DBPort        string `env:"DBPORT"`
	DBUser        string `env:"DBUSER"`
	DBPassword    string `env:"DBPASSWORD"`
	NeedAuthen    bool   `env:"NEEDAUTHEN"`
	JWTSecret     string `env:"JWTSECRET"`
	RedisAddr     string `env:"REDISADDR"`
	RedisPassword string `env:"REDISPASSWORD"`
	RedisDB       int    `env:"REDISDB"`
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
