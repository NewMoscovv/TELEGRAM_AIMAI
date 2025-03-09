package config

import (
	"AIMAI/pkg/consts"
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	BotSettings BotSettings
}

type BotSettings struct {
	TelegramToken string
}

func Init() (*Config, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		return nil, errors.New(consts.TelegramTokenIsAbsent)
	}

	return &Config{
		BotSettings: BotSettings{TelegramToken: telegramToken},
	}, nil

}
