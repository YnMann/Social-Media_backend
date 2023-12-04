package api_server

// Config
type Config struct {
	// Адрес по которому мы запускаем веб сервер
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level`
}

// Передает инициализированный config, default параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
