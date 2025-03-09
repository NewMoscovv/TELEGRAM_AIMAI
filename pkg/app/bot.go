package app

import (
	"AIMAI/pkg/config"
	myLogger "AIMAI/pkg/logger"
	"AIMAI/pkg/middleware"
	tele "gopkg.in/telebot.v3"
	"time"
)

type BotConfig struct {
	Self   *tele.Bot
	logger *myLogger.Logger
}

func NewBot(cfg config.BotSettings, logger *myLogger.Logger) (*BotConfig, error) {

	// настройка характеристик бота
	pref := tele.Settings{
		Token:     cfg.TelegramToken,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: tele.ModeMarkdown,
	}

	// создание экземлпяра бота
	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return &BotConfig{Self: bot, logger: logger}, nil

}

func (bot *BotConfig) SetupHandlers() {

	bot.Self.Handle("/start", middleware.LoggingMiddleware(func(c tele.Context) error {
		return c.Send("hello world")
	}, bot.logger))
}

func (bot *BotConfig) Start() {

	bot.Self.Start()

}
