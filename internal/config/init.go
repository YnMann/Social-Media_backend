import "github.com/YnMann/chat_backend/internal/server/server"

// Config
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// Передает инициализированный config, default параметрами
func ConfigInit() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}