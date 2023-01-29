package messages

type Message struct {
	Text   string
	UserID int64
}

type MessageSender interface {
	SendMessage(Text string, UserID int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{tgClient: tgClient}

}

func (m *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {
		m.tgClient.SendMessage("hello", msg.UserID)
		return nil
	}
	m.tgClient.SendMessage("unknown command", msg.UserID)
	return nil

}
