package store

type Config struct {
	DatabaseDriver     string `toml:"db_driver"`
	DatabaseConnection string `toml:"db_connection"`
}

func NewConfig() *Config {
	return &Config{}
}
