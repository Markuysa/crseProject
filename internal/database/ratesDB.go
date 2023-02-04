package database

import (
	"context"
	"log"
	"ozonProjectmodule/internal/model/domain"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type RatesDB struct {
	db *sqlx.DB
}

func NewRatesDB(db *sqlx.DB) *RatesDB {
	return &RatesDB{db: db}
}

func (db *RatesDB) Add(ctx context.Context, rate domain.Rate) error {
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

	builder := sq.Select(
		"id",
		"code",
		"nominal",
		"kopecks",
		"original",
		"ts",
		"created_at",
		"updated_at",
		"deleted_at",
	).From("rates").Where(sq.Eq{"code": code})

	if !date.IsZero() {
		builder = builder.Where(sq.Eq{"ts": date})
	}
	query, args, err := builder.ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "conversion to sql code in builder")
	}
	log.Print(builder)
	var rate domain.Rate

	err = db.db.QueryRowContext(ctx, query, args...).Scan(&rate)
	if err != nil {
		return nil, errors.Wrap(err, "queryrow in database")
	}
	return &rate, nil
}
