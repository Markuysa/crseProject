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

	switch msg.Text {
	case "/start":
		m.tgClient.SendMessage("hello", msg.UserID)
	case "/add":
		m.tgClient.SendMessage("Enter the amount:", msg.UserID)
		//service.Add()
	case "/weekreport":
		m.tgClient.SendMessage("Form the weekly report...", msg.UserID)

	case "/monthreport":
		m.tgClient.SendMessage("Form the monthly report...", msg.UserID)
	case "/yearreport":
		m.tgClient.SendMessage("Form the yearly report...", msg.UserID)
	default:
		m.tgClient.SendMessage("Unknown command", msg.UserID)
	}
	return nil

}
