package app

import (
	"AIMAI/pkg/config"
	tele "gopkg.in/telebot.v3"
	"time"
)

type BotConfig struct {
	Self *tele.Bot
}

func NewBot(cfg config.BotSettings) (*BotConfig, error) {

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

	return &BotConfig{Self: bot}, nil

}

func (bot *BotConfig) SetupHandlers() {

	bot.Self.Handle("/start", func(c tele.Context) error {
		return c.Send("hello world")
	})
}

func (bot *BotConfig) Start() {

	bot.Self.Start()

}
