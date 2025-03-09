package app

import (
	"AIMAI/internal/middleware"
	"AIMAI/internal/openrouter"
	"AIMAI/internal/user"
	"AIMAI/pkg/config"
	myLogger "AIMAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
	"sync"
	"time"
)

type BotConfig struct {
	Self       *tele.Bot
	Middleware *middleware.Middleware
	Messages   config.Messages
	Users      map[int64]user.User
	OpenRtr    *openrouter.Client
	mu         sync.RWMutex
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

	// инициализируем openrouter
	openRtr := openrouter.NewClient(cfg.OpenRtr.APIKey, cfg.OpenRtr.APIUrl, cfg.OpenRtr.Model)

	// создание мапы юзерс
	users := make(map[int64]user.User)

	return &BotConfig{
		Self:       bot,
		Middleware: middlewares,
		Messages:   cfg.Messages,
		OpenRtr:    openRtr,
		Users:      users}, nil

}

func (bot *BotConfig) SetupHandlers() {

	bot.Self.Handle("/start", bot.Middleware.LoggingMiddleware(bot.HandleStart))
	bot.Self.Handle(tele.OnText, bot.Middleware.LoggingMiddleware(bot.HandlerMessage))
}

func (bot *BotConfig) Start() {

	bot.Self.Start()

}

func (bot *BotConfig) getUser(chatID int64) (user.User, bool) {
	bot.mu.RLock()
	defer bot.mu.RUnlock()
	curUser, ok := bot.Users[chatID]
	return curUser, ok
}

func (bot *BotConfig) setUserStatus(chatID int64, user user.User) {
	bot.mu.Lock()
	defer bot.mu.Unlock()
	bot.Users[chatID] = user
}
