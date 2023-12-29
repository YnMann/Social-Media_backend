package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	server "github.com/YnMann/chat_backend/internal/server/server"
)

var (
	config_path string
)

func init() {
	flag.StringVar(
		&config_path,
		"config-path",
		"config/api_server.toml",
		"path to config file")
}

func main() {
	flag.Parse()
	// Создание нового объекта конфигурации API-сервера
	config := server.NewConfig()
	_, err := toml.DecodeFile(config_path, config)
	if err != nil {
		log.Fatal(err)
	}
	// Создание экземпляра API-сервера
	s := server.New(config)

	// Если имеется ошибка при создании сервера,
	// то выбрасываем ошибку
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
