package main

import (
	myApp "AIMAI/pkg/app"
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
	app, err := myApp.NewApp(cfg)
	if err != nil {
		logger.Err.Panic(err)
	}

	//запуск
	app.Start()

}
