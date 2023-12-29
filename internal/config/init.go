package config

import "github.com/spf13/viper"

// Config
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	// Store    *store.Config
}

func Init() (*Config, error) {
	// Инициализируем Viper
	viper.SetConfigName("config")
	viper.AddConfigPath("../../config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Создаем структуру Config и заполняем ее значениями из Viper
	config := &Config{
		BindAddr: viper.GetString("bind_addr"),
		LogLevel: viper.GetString("log_level"),
		// Store: &store.Config{
		// 	DatabaseURL: viper.GetString("store.database_url"),
		// },
	}

	return config, nil
}
