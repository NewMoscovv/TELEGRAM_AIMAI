package app

import (
	"AIMAI/pkg/config"
	"AIMAI/pkg/logger"
)

type App struct {
	Bot *BotConfig
}

func NewApp(cfg *config.Config, logger *logger.Logger) (*App, error) {

	bot, err := NewBot(cfg.BotSettings, logger)
	if err != nil {
		return nil, err
	}

	return &App{
		Bot: bot,
	}, nil

}

func (app *App) Start() {

	app.Bot.SetupHandlers()

	app.Bot.Start()
}
