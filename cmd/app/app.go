package main

import (
	"fmt"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db"
	"github.com/go-telegram/bot"
)

type application struct {
	bot     *bot.Bot
	storage *db.Storage
}

func (app *application) run() {
	fmt.Println("Telegram bot started")
}
