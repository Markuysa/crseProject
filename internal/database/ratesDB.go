package database

import (
	"context"
	"ozonProjectmodule/internal/model/domain"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type RatesDB struct {
	db *sqlx.DB
}

func NewRatesDB(db *sqlx.DB) *RatesDB {
	return &RatesDB{db: db}
}

func (db *RatesDB) AddRate(ctx context.Context, rate domain.Rate) error {
	var query = `
	insert into rates(
		created_at,
		code,
		nominal,
		kopecks,
		original,
		ts
	)values(
		$1,$2,$3,$4,$5,$6
	);
	`
	_, err := db.db.ExecContext(ctx, query,
		rate.CreatedAt,
		rate.Code,
		rate.Nominal,
		rate.Kopecks,
		rate.Original,
		rate.Ts,
	)
	if err != nil {
		return errors.Wrap(err, "Adding new rate:")
	}
	return nil
}

func (db *RatesDB) GetRate(ctx context.Context, code string, date time.Time) (*domain.Rate, error) {
	query := `
	select 	id,
			code,
			nominal,
			kopecks,
			original,
			ts,
			created_at,
			updated_at,
			deleted_at
	from rates
	where code = $1 AND ts = $2
	`
	var rate domain.Rate

	err := db.db.QueryRowContext(ctx, query, code, date).Scan(&rate.ID, &rate.Code, &rate.Nominal,
		&rate.Kopecks, &rate.Original, &rate.Ts, &rate.CreatedAt, &rate.UpdatedAt, &rate.DeletedAt)
	if err != nil {
		return nil, errors.Wrap(err, "queryrow in database")
	}
	return &rate, nil
}
