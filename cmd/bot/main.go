package main

import (
	"log"
	"ozonProjectmodule/internal/clients/tg"
	"ozonProjectmodule/internal/config"
	"ozonProjectmodule/internal/model/messages"
)

func main() {

	config, err := config.New()

	if err != nil {
		log.Fatal("config init failed", err)
	}

	tgClient, err := tg.New(config)

	if err != nil {
		log.Fatal("tg client init failed", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)

}
