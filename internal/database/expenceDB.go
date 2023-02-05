package database

import (
	"context"
	"log"
	"ozonProjectmodule/internal/model/domain"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Expenditure_type string
	Date             time.Time
	Amount           int64 // в копейках
}

type ExpenseDB struct {
	db *sqlx.DB
}

func NewExpenseDB(db *sqlx.DB) *ExpenseDB {

	return &ExpenseDB{
		db: db,
	}

}

func (db *ExpenseDB) AddExpence(ctx context.Context, expenditure domain.Expenditure) error {

	query := `
		insert into expenditures (
			anmount,
			expenditure_type,
			date
		)values(
			$1,$2,$3
		);

	`
	_, err := db.db.ExecContext(ctx, query, expenditure.Anmount, expenditure.ExpenditureType,
		expenditure.Date)
	if err != nil {
		return errors.Wrap(err, "adding expense")
	}
	return nil

}
func (db *ExpenseDB) GetExpense(ctx context.Context, date time.Time) (*Expense, error) {

	var expence Expense

	query := `
		select 	anmount,
				expenditure_type,
				date
		from expenditures
		where date = $1
	`

	result, err := db.db.QueryxContext(ctx, query, date)
	log.Println(date)
	if err != nil {
		return nil, errors.Wrap(err, "query expese")
	}
	result.StructScan(&expence)

	return &expence, nil

}
