package main

import (
	"log"

	"github.com/YnMann/chat_backend/internal/config"
	"github.com/YnMann/chat_backend/internal/server"
	"github.com/spf13/viper"
)

func main() {
	// Создание нового объекта конфигурации API-сервера
	if _, err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	// Создание экземпляра API-сервера
	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
