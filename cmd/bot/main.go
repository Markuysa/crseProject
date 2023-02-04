package main

import (
	"context"
	"log"
	"ozonProjectmodule/internal/clients/tg"
	"ozonProjectmodule/internal/config"
	"ozonProjectmodule/internal/database"
	"ozonProjectmodule/internal/model/domain"
	"ozonProjectmodule/internal/model/messages"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

func main() {

	config, err := config.New()

	if err != nil {
		log.Fatal("config init failed", err)
	}

	//-----------------------------------------------------------------------------------------------//
	ctx := context.Background()

	db, err := sqlx.Open("postgres", "host=localhost port=5431 user=postgres password=pass sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	ratesDB := database.NewRatesDB(db)
	object := domain.Rate{
		Code:     "RUB",
		Nominal:  12,
		Kopecks:  12,
		Original: "Rubles",
		Ts:       time.Now(),
	}
	ratesDB.Add(ctx, object)
	log.Print(ratesDB.GetRate(ctx, object.Code, object.Ts))
	tgClient, err := tg.New(config)

	if err != nil {
		log.Fatal("tg client init failed", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}
