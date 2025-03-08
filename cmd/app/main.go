package main

import (
	"log"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/env"
	"github.com/go-telegram/bot"
)

func main() {
	b, err := bot.New(env.Get("BOT_TOKEN"))
	if nil != err {
		log.Fatalf("Failed to create bot: %v", err)
	}

	storage := db.NewStorage()

	app := &application{bot: b, storage: storage}

	app.run()
}
