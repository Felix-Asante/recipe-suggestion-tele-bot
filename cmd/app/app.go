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
	app.registerHandlers()

}

func (app *application) registerHandlers() {
	app.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, app.startHandler)
	app.bot.RegisterHandler(bot.HandlerTypeMessageText, "/find", bot.MatchTypeExact, app.findRecipesHandler)
	app.bot.RegisterHandler(bot.HandlerTypeMessageText, "/saved", bot.MatchTypeExact, app.savedRecipesHandler)
	app.bot.RegisterHandler(bot.HandlerTypeMessageText, "/diet", bot.MatchTypeExact, app.dietHandler)
	// app.bot.RegisterHandler(bot.HandlerTypeMessageText, "/mealplan", bot.MatchTypeExact, app.mealPlanHandler)
	app.bot.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, app.helpHandler)
}
