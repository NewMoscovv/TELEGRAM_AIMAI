package user

import (
	"AIMAI/internal/openrouter"
	"AIMAI/pkg/consts"
	"errors"
)

type User struct {
	Status      Status
	ChatHistory []openrouter.Message
}

type Status string

func (u *User) AddMessage(text string, role string) error {

	if len(u.ChatHistory) <= consts.MaxFreeDialogLen {
		u.ChatHistory = append(u.ChatHistory, openrouter.Message{
			Role:    role,
			Content: text,
		})
		return nil
	}

	return errors.New("превышен лимит сообщений")

}

func (u *User) GetChatHistory() []openrouter.Message {
	return u.ChatHistory
}

func (u *User) ResetHistory() {
	u.ChatHistory = []openrouter.Message{}
}
