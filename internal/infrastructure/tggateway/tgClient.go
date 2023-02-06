package tg

import (
	"context"
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

func (c *Client) ListenUpdates(ctx context.Context, msgModel *messages.Model) {

	updates := c.Start()

	for update := range updates {
		if update.Message != nil {

			err := msgModel.IncomingMessage(ctx, messages.Message{
				Text:   update.Message.Text,
				UserID: update.Message.From.ID,
			})
			if err != nil {
				log.Println("error processing message", err)
			}

		}
	}

}
func (c *Client) Start() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return c.client.GetUpdatesChan(u)
}

func (c *Client) Stop() {
	c.client.StopReceivingUpdates()
}

func (c *Client) Request(callback tgbotapi.CallbackConfig) error {
	_, err := c.client.Request(callback)
	return err
}
