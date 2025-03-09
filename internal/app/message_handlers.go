package app

import tele "gopkg.in/telebot.v3"

func (bot *BotConfig) HandlerStart(c tele.Context) error {
	return c.Send(bot.Messages.Responses.WelcomeMsg)
}
