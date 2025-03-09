package config

import (
	"AIMAI/pkg/consts"
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	BotSettings BotSettings
}

type BotSettings struct {
	TelegramToken string
	Messages      Messages
}

type Messages struct {
	Responses Responses
}

type Responses struct {
	WelcomeMsg string `mapstructure:"welcome_message"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.responses", &config.BotSettings.Messages.Responses); err != nil {
		return nil, err
	}

	err := parseEnv(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func parseEnv(cfg *Config) error {

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		return errors.New(consts.TelegramTokenIsAbsent)
	}

	cfg.BotSettings.TelegramToken = telegramToken

	return nil
}
