package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"

	"semen_project/internal/pkg/pgconn"
)

type Config struct {
	AppName string `env:"APP_NAME,notEmpty" envDefault:"semen-platform"`

	PublicApiHost string `env:"PUBLIC_API_HOST"`       
	PublicApiPort int `env:"PUBLIC_API_PORT,notEmpty" envDefault:"8080"`
	
	JWTSecret string `env:"JWT_SECRET,notEmpty"`

	PG pgconn.DbConfig
}

func  Load(cfg *Config) error {
	err := env.Parse(cfg)
	if err != nil {
		return fmt.Errorf("Config.Load(): %w", err)
	}

	return nil
}

func (cfg *Config) PublicApiAddr() string {        
	return fmt.Sprintf("%s:%d", cfg.PublicApiHost, cfg.PublicApiPort)
}