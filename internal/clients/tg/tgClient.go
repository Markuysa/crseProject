package tg

import (
	"log"
	"ozonProjectmodule/internal/model/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type TokenGetter interface {
	Token() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tokenGetter TokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())

	if err != nil {
		return nil, errors.Wrap(err, "NewBotApi")
	}
	return &Client{
		client: client,
	}, nil
}
func (c *Client) SendMessage(text string, userID int64) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userID, text))
	if err != nil {
		return errors.Wrap(err, "send message")
	}
	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message, update.Message.ForwardFrom.ID)

			err := msgModel.IncomingMessage(messages.Message{
				Text:   update.Message.Text,
				UserID: update.Message.From.ID,
			})
			if err != nil {
				log.Println("error processing message", err)
			}

		}
	}

}