package messages

import (
	"context"
	"ozonProjectmodule/internal/model/domain"
	"strings"
)

const newUser = `Привет! Я буду помогать ввести твою бухгалтерию. Но перед началом работы тебе нужны выбрать валюту по умолчанию в которой ты производишь расходы`

const introMessage = "Привет! Я умею учитывать твои траты.\n\n" + helpMessage

const helpMessage = `Для работы с ботом тебе могут потребоваться следующие команды:
Чтобы изменить выбранную валюту необходимо выполнить команду /change_currency

Чтобы добавить новый расход, отправь мне сообщение в формате:
/add цена; описание; дата (дд.мм.гггг, опционально) - если не указать дату, то расход будет добавлен на сегодняшний день

Чтобы посмотреть расходы отправь:
/list -  за всё время.
/list_day - за день.
/list_week - за неделю.
/list_month - за месяц.
/list_year - за год.`

const unknownMessage = `Неизвестная команда. Чтобы посмотреть список команд отправь /help`

func (m *Model) IncomingMessage(ctx context.Context, msg Message) error {

	// switch msg.Text {
	// case "/start":
	// 	m.tgClient.SendMessage(introMessage, msg.UserID)
	// case "/help":
	// 	m.tgClient.SendMessage(helpMessage, msg.UserID)
	// case "/add":
	// 	m.tgClient.SendMessage("Enter the amount:", msg.UserID)
	// 	//service.Add()
	// case "/weekreport":
	// 	m.tgClient.SendMessage("Form the weekly report...", msg.UserID)

	// case "/monthreport":
	// 	m.tgClient.SendMessage("Form the monthly report...", msg.UserID)
	// case "/yearreport":
	// 	m.tgClient.SendMessage("Form the yearly report...", msg.UserID)
	// default:
	// 	m.tgClient.SendMessage(unknownMessage, msg.UserID)
	// }

	if strings.HasPrefix(msg.Text,"/set_currency") && !m.userDB.UserExist(ctx, msg.UserID) {
		m.userDB.AddUser(ctx,domain.User{
			UserID: msg.UserID,
			DefaultCurrency: strings.Trim(msg.Text,"/set_currency"),
		})
	}

	return nil

}
