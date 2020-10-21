package apiserver

import "github.com/EDEN45/go-blog/internal/app/store"

type Config struct {
	AppPort  string `toml:"app_port"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		AppPort:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
