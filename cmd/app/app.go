package main

import (
	"fmt"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/repositories"
	"github.com/go-telegram/bot"
)

type application struct {
	bot          *bot.Bot
	repositories *repositories.Repositories
}

func (app *application) run() {
	fmt.Println("Telegram bot started")

	app.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, app.startHandler)
}
