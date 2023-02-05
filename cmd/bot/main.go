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

	db2 := database.NewRatesDB(db)
	database.NewExpenseDB(db)
	database.NewUserDB(db)
	rate := domain.Rate{
		CreatedAt: time.Now(),
		Code:      "12",
		Nominal:   12,
		Kopecks:   12,
		Original:  "as",
		Ts:        time.Now(),
	}
	db2.AddRate(ctx, rate)
	log.Println(db2.GetRate(ctx, rate.Code, rate.CreatedAt))
	tgClient, err := tg.New(config)

	if err != nil {
		log.Fatal("tg client init failed", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)

}
