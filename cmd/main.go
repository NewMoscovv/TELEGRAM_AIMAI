package main

import (
	myApp "AIMAI/internal/app"
	"AIMAI/pkg/config"
	myLogger "AIMAI/pkg/logger"
)

func main() {

	// инициируем логгер
	logger := myLogger.Init()

	// инициируем конфиг
	cfg, err := config.Init()
	if err != nil {
		logger.Err.Panic(err)
	}

	// инициируем приложение
	app, err := myApp.NewApp(cfg, logger)
	if err != nil {
		logger.Err.Panic(err)
	}

	//запуск
	app.Start()

}
