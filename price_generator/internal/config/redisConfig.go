// Package config to env
package config

import "github.com/caarlos0/env/v6"

// RedisConfig struct to redis config env
type RedisConfig struct {
	Addr     string `env:"ADDR_REDIS" envDefault:"redis:6379"`
	Password string `env:"PASSWORD_REDIS" envDefault:""`
	DB       int    `env:"DB_REDIS" envDefault:"0"`
}

// NewRedis contract redis config
func NewRedis() (*RedisConfig, error) {
	cfg := &RedisConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
