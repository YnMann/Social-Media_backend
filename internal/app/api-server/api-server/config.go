package api_server

import "github.com/YnMann/social-media_backend/internal/app/api-server/store"

// Config
type Config struct {
	// Адрес по которому мы запускаем веб сервер
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// Передает инициализированный config, default параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
