package app

import (
	"AIMAI/internal/user/states"
	"AIMAI/pkg/consts"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"time"
)

func (bot *BotConfig) HandleStart(c tele.Context) error {

	_, ok := bot.getUser(c.Sender().ID)
	if !ok {
		bot.createUser(c.Sender().ID)
	}

	return c.Send(bot.Messages.Responses.WelcomeMsg)
}

func (bot *BotConfig) HandlerMessage(c tele.Context) error {

	status := bot.getUserStatus(c.Sender().ID)
	if status == states.WaitingForResponse {

		return c.Send("Дождитесь генерации сообщения")
	}

	err := bot.addMessage(c.Text(), c.Sender().ID, "user")

	if err != nil {
		inlineKeys := [][]tele.InlineButton{
			{tele.InlineButton{Text: "Сбросить историю",
				Unique: "reset_dialog_history"}},
		}
		return c.Send("Превышен лимит сообщений...", &tele.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		})
	}

	bot.setUserStatus(c.Sender().ID, states.WaitingForResponse)

	for i := 0; i < consts.MaxAmountResponses; i++ {
		// Печатает...
		c.Notify(tele.Typing)

		chHistory := bot.getUserChatHistory(c.Sender().ID)

		response, err := bot.OpenRtr.GetResponse(chHistory)
		bot.addMessage(response, c.Sender().ID, "assistant")

		fmt.Println(response)

		if err != nil {
			bot.logger.Err.Println(err)
			return c.Send("Ой, что-то пошло не так. Обратитесь в поддержку - <b>@new_moscovv</b>")
		}
		if response == "" {
			bot.logger.Info.Printf("пустой ответ от ИИ, выполнение повторного запроса...")
		} else {
			bot.setUserStatus(c.Sender().ID, states.Default)
			return c.Send(response)
		}
		time.Sleep(1 * time.Second)
	}

	return c.Send("Произошла ошибка при получении ответа от ИИ. Пожалуйста, повторите запрос или обратитесь в поддержку")
}

func (bot *BotConfig) HandleCallback(c tele.Context) error {

	c.Respond(&tele.CallbackResponse{Text: "история диалога сброшена"})

	c.Delete()

	user, _ := bot.getUser(c.Sender().ID)
	user.ResetHistory()
	bot.Users[c.Sender().ID] = user

	return c.Send("диалог подчищен черт")
}
