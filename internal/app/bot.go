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

	// инициализация мидлвари
	middlewares := middleware.NewMiddleware(logger)

	// инициализируем openrouter
	openRtr := openrouter.NewClient(cfg.OpenRtr.APIKey, cfg.OpenRtr.APIUrl, cfg.OpenRtr.Model, cfg.OpenRtr.Prompt)

	// создание мапы юзерс
	users := make(map[int64]user.User)

	return &BotConfig{
		Self:       bot,
		Middleware: middlewares,
		Messages:   cfg.Messages,
		OpenRtr:    openRtr,
		Users:      users,
		logger:     logger,
	}, nil

}

func (bot *BotConfig) SetupHandlers() {

	bot.Self.Handle("/start", bot.HandleStart, bot.Middleware.LoggingMiddleware)
	bot.Self.Handle(tele.OnText, bot.HandlerMessage, bot.Middleware.LoggingMiddleware)
	bot.Self.Handle(&tele.InlineButton{Unique: "reset_dialog_history"}, bot.HandleCallback)
}

func (bot *BotConfig) Start() {

	bot.Self.Start()

}

func (bot *BotConfig) createUser(chatID int64) {
	bot.mu.Lock()
	defer bot.mu.Unlock()

	history := make([]openrouter.Message, 0)

	bot.Users[chatID] = user.User{
		Status:      "default",
		ChatHistory: history,
	}

}

func (bot *BotConfig) getUser(chatID int64) (user.User, bool) {
	bot.mu.RLock()
	defer bot.mu.RUnlock()
	curUser, ok := bot.Users[chatID]
	return curUser, ok
}

func (bot *BotConfig) getUserStatus(chatID int64) user.Status {
	bot.mu.RLock()
	defer bot.mu.RUnlock()
	curUser := bot.Users[chatID]

	return curUser.Status
}

func (bot *BotConfig) setUserStatus(chatID int64, status user.Status) {
	bot.mu.Lock()
	defer bot.mu.Unlock()
	user := bot.Users[chatID]
	user.Status = status
	bot.Users[chatID] = user
}

func (bot *BotConfig) getUserChatHistory(chatID int64) []openrouter.Message {
	curUser, _ := bot.getUser(chatID)
	return curUser.GetChatHistory()
}

func (bot *BotConfig) addMessage(message string, chatID int64, role string) error {
	curUser, _ := bot.getUser(chatID)
	err := curUser.AddMessage(message, role)

	if err != nil {
		return err
	}

	bot.mu.Lock()
	defer bot.mu.Unlock()
	bot.Users[chatID] = curUser

	return nil
}
