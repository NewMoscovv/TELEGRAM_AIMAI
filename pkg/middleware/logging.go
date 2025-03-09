package middleware

import (
	myLogger "AIMAI/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

func LoggingMiddleware(next tele.HandlerFunc, logger *myLogger.Logger) tele.HandlerFunc {
	return func(c tele.Context) error {
		username := "unknown"
		if c.Sender().Username != "" {
			username = c.Sender().Username
		}
		logger.Info.Printf("%s | %s", username, c.Message().Text)

		return next(c)
	}

}
