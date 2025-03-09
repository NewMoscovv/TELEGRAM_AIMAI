package config

import (
	"AIMAI/internal/openrouter"
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

	OpenRtr openrouter.Client
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

	openRouterToken := os.Getenv("OPENROUTER_TOKEN")
	if openRouterToken == "" {
		return errors.New(consts.OpenRouterTokenIsAbsent)
	}

	apiUrl := os.Getenv("API_URL")
	if apiUrl == "" {
		return errors.New(consts.OpenRouterUrlIsAbsent)
	}

	model := os.Getenv("MODEL")
	if model == "" {
		return errors.New(consts.OpenRouterModelIsAbsent)
	}

	prompt := os.Getenv("PROMPT")
	if prompt == "" {
		return errors.New(consts.OpenRouterPromptIsAbsent)
	}

	cfg.BotSettings.OpenRtr = openrouter.Client{
		APIUrl: apiUrl,
		Model:  model,
		APIKey: openRouterToken,
		Prompt: prompt,
	}

	cfg.BotSettings.TelegramToken = telegramToken

	return nil
}
