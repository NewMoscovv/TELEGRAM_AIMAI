package middleware

import (
	myLogger "AIMAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

type Middleware struct {
	logger *myLogger.Logger
}

func NewMiddleware(logger *myLogger.Logger) *Middleware {
	return &Middleware{logger: logger}
}

func (m *Middleware) LoggingMiddleware(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		username := "unknown"
		if c.Sender().Username != "" {
			username = c.Sender().Username
		}
		m.logger.Info.Printf("%s | %s", username, c.Message().Text)

		err := next(c)

		if err != nil {
			m.logger.Err.Printf("ТЕЛО ОШИБКИ - %s", err)
			return err
		}
		return nil
	}
}
