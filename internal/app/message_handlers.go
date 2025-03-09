package app

import (
	"AIMAI/pkg/consts"
	tele "gopkg.in/telebot.v3"
	"log"
	"strings"
	"time"
)

func (bot *BotConfig) HandleStart(c tele.Context) error {
	return c.Send(bot.Messages.Responses.WelcomeMsg)
}

func (bot *BotConfig) HandlerMessage(c tele.Context) error {

	for i := 0; i < consts.MaxAmountResponses; i++ {
		// Печатает...
		c.Notify(tele.Typing)

		response, err := bot.OpenRtr.GetResponse(c.Text())
		if err != nil {
			return c.Send("Ой, что-то пошло не так. Обратитесь в поддержку - <b>@new_moscovv</b>")
		}
		if response == "" {
			log.Printf("пустой ответ от ИИ, выполнение повторного запроса...")
		} else {
			log.Printf("%s | %s", bot.Self.Me.Username, strings.Replace(response, "\n\n", "\n", -1))
			return c.Send(response)
		}
		time.Sleep(1 * time.Second)
	}

	return c.Send("Произошла ошибка при получении ответа от ИИ. Пожалуйста, повторите запрос или обратитесь в поддержку")
}
