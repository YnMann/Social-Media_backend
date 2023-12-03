package api_server

// Config
type Config struct {
	// Адрес по которому мы запускаем веб сервер
	BindAddr string `toml:"bind_addr"`
}

// Передает инициализированный config, default параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
