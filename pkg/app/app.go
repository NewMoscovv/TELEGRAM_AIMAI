package app

import "AIMAI/pkg/config"

type App struct {
	Bot *BotConfig
}

func NewApp(cfg *config.Config) (*App, error) {

	bot, err := NewBot(cfg.BotSettings)
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
