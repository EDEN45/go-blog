package apiserver

type Config struct {
	AppPort  string `toml:"app_port"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		AppPort:  ":8080",
		LogLevel: "debug",
	}
}
