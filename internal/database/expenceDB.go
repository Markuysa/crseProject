package database

import (
	"context"
	"ozonProjectmodule/internal/model/domain"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ExpenseDB struct {
	db *sqlx.DB
}

func NewExpenseDB(db *sqlx.DB) *ExpenseDB {

	return &ExpenseDB{
		db: db,
	}

}

func (db *ExpenseDB) AddExpence(ctx context.Context, expenditure domain.Expense) error {

	query := `
		insert into expenditures (
			anmount,
			expenditure_type,
			date
		)values(
			$1,$2,$3
		);

	`
	_, err := db.db.ExecContext(ctx, query, expenditure.Amount, expenditure.Expenditure_type,
		expenditure.Date)
	if err != nil {
		return errors.Wrap(err, "adding expense")
	}
	return nil

}
func (db *ExpenseDB) GetExpenses(ctx context.Context, userID int64) ([]domain.Expense, error) {

	var expenses []domain.Expense

	query := `
		select 	anmount,
				expenditure_type,
				date
		from expenditures
		where userID = $1
	`

	err := db.db.SelectContext(ctx, &expenses, query, userID)
	if err != nil {
		return nil, errors.Wrap(err, "query expese")
	}

	return expenses, nil

}
