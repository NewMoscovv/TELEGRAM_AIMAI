package logger

import (
	"log"
	"os"
)

type Logger struct {
	Err  *log.Logger
	Info *log.Logger
}

func InitLogger() *Logger {
	logger := Logger{}
	logger.Info = log.New(os.Stdout, "[ИНФО]   ", log.Ldate|log.Ltime)
	logger.Err = log.New(os.Stderr, "[ОШИБКА] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info.Println("Логгер инициализирован")
	return &logger
}
