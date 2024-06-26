package telegram

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	T *tg.BotAPI
}

func NewBot(APIToken string) (*Bot, error) {
	t, err := tg.NewBotAPI(APIToken)
	if err != nil {
		return &Bot{}, err
	}

	t.Debug = false

	fmt.Printf("Authorized on account %s\n", t.Self.UserName)
	return &Bot{t}, nil
}

func (b *Bot) SendVideo(filepath string, channelID int64) (tg.Message, error) {
	c := tg.NewVideoUpload(channelID, filepath)
	c.BaseChat.DisableNotification = true

	fmt.Println("uploading", filepath)
	msg, err := b.T.Send(c)
	if err != nil {
		return tg.Message{}, err
	}

	return msg, nil
}
