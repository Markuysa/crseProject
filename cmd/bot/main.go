package main

import (
	"context"
	"log"
	"ozonProjectmodule/internal/config"
	"ozonProjectmodule/internal/database"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/jmoiron/sqlx"
)

func main() {
	ctx := context.Background()
	config, err := config.New()

	if err != nil {
		log.Fatal("config init failed", err)
	}

	//-----------------------------------------------------------------------------------------------//
	//ctx := context.Background()

	db, err := sqlx.Open("postgres", "host=localhost port=5431 user=postgres password=pass sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	dbUser := database.NewUserDB(db)
	dbRates := database.NewRatesDB(db)
	dbExpense := database.NewExpenseDB(db)


	

}
