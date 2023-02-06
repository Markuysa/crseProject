package messages

import (
	"context"
	"ozonProjectmodule/internal/model/domain"
	"time"
)

type Message struct {
	Text   string
	UserID int64
}

type MessageSender interface {
	SendMessage(Text string, UserID int64) error
}

type ExpenseDB interface {
	AddExpence(ctx context.Context, expenditure domain.Expense) error
	GetExpenses(ctx context.Context, userID int64) ([]domain.Expense, error)
}

type UserDB interface {
	GetUser(ctx context.Context, userId int64) (*domain.User, error)
	AddUser(ctx context.Context, user domain.User) error
	UserExist(ctx context.Context, userID int64) bool
	ChangeDefaultCurrency(ctx context.Context, userID int64, currency string) error
	GetDefaultCurrency(ctx context.Context, userID int64) (string, error)
}

type RateDB interface {
	GetRate(ctx context.Context, code string, date time.Time) (*domain.Rate, error)
	AddRate(ctx context.Context, rate domain.Rate) error
}
type Model struct {
	tgClient    MessageSender
	expenseDB   ExpenseDB
	userDB      UserDB
	rateDB      RateDB
	config      ConfigGetter
	rateUpdater CurrencyExchangeRateUpdater
}
type ConfigGetter interface {
	SupportedCurrencyCodes() []string
	GetBaseCurrency() string
}
type CurrencyExchangeRateUpdater interface {
	UpdateCurrencyExchangeRatesOn(ctx context.Context, time time.Time) error
}

func New(tgClient MessageSender) *Model {
	return &Model{tgClient: tgClient}

}
