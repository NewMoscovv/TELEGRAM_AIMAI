package app

import (
	"AIMAI/internal/middleware"
	"AIMAI/pkg/config"
	myLogger "AIMAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
	"time"
)

type BotConfig struct {
	Self       *tele.Bot
	Middleware *middleware.Middleware
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

	// инициализация мидлвари
	middlewares := middleware.NewMiddleware(logger)

	return &BotConfig{Self: bot, Middleware: middlewares}, nil

}

func (bot *BotConfig) SetupHandlers() {

	bot.Self.Handle("/start", bot.Middleware.LoggingMiddleware(func(c tele.Context) error {
		return c.Send("hello world")
	}))
}

func (bot *BotConfig) Start() {

	bot.Self.Start()

}
