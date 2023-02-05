package database

import (
	"context"
	"ozonProjectmodule/internal/model/domain"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID          int64
	DefaultCurrency string
}

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {

	return &UserDB{
		db: db,
	}

}
func (db *UserDB) AddUser(ctx context.Context, user domain.User) error {

	query := `
		insert into users (
			user_id,
			default_currency
		)values(
			$1,$2
		);

	`
	_, err := db.db.ExecContext(ctx, query, user.UserID, user.DefaultCurrency)
	if err != nil {
		return errors.Wrap(err, "adding expense")
	}
	return nil

}
func (db *UserDB) GetUser(ctx context.Context, userId int64) (*domain.User, error) {

	var user domain.User

	query := `
		select 	user_id,
				default_currency
		from users
		where user_id = 1234
	`

	result, err := db.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "query expese")
	}
	result.StructScan(&user)

	return &user, nil

}
