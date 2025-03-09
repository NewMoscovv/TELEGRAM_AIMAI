package app

import (
	tele "gopkg.in/telebot.v3"
)

func (bot *BotConfig) HandlerStart(c tele.Context) error {

	user, ok := bot.getUser(c.Sender().ID)
	if ok {
		return c.Send(string(user.Status))
	}
	user.Status = "default"
	bot.setUserStatus(c.Sender().ID, user)

	return c.Send(bot.Messages.Responses.WelcomeMsg)
}
