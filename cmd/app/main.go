package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	api_server "github.com/YnMann/social-media_backend/internal/app/api-server/api-server"
)

var (
	config_path string
)

func init() {
	flag.StringVar(
		&config_path,
		"config-path",
		"configs/apis",
		"path to config file")
}

func main() {
	flag.Parse()
	// Создание нового объекта конфигурации API-сервера
	config := api_server.NewConfig()
	_, err := toml.DecodeFile(config_path, config)
	if err != nil {
		log.Fatal(err)
	}

	// Создание экземпляра API-сервера
	s := api_server.New(config)

	// Если имеется ошибка при создании сервера,
	// то выбрасываем ошибку
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
