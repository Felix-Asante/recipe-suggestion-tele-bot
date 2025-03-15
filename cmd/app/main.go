package main

import (
	"context"
	"log"
	"net/http"

	"os"
	"os/signal"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/repositories"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/env"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/utils"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var app *application

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	defer utils.RecoverFromPanic()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(env.GetString("BOT_TOKEN", ""), opts...)
	if nil != err {
		panic(err)
	}

	b.SetWebhook(ctx, &bot.SetWebhookParams{
		URL: env.GetString("APP_URL", ""),
	})

	if _, err := setCommands(ctx, b); nil != err {
		log.Fatal(err)
	}

	var appErr error

	app, appErr = createApp(b)

	if nil != appErr {
		log.Fatal(appErr)
	}

	app.run()

	go func() {
		http.ListenAndServe(env.GetString("PORT", ":2000"), b.WebhookHandler())
	}()

	b.StartWebhook(ctx)

}

func createApp(b *bot.Bot) (*application, error) {
	db, err := db.New()
	if nil != err {
		return nil, err
	}

	repositories := repositories.NewRepositories(db)

	app := &application{bot: b, repositories: repositories}

	return app, nil
}

func setCommands(ctx context.Context, b *bot.Bot) (bool, error) {
	return b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{
				Command:     "/start",
				Description: "Show welcome message",
			},
			{
				Command:     "/findrecipe",
				Description: "Search recipes with your ingredients",
			},
			// {
			// 	Command:     "/pantry",
			// 	Description: "Manage your saved ingredient list",
			// },
			{
				Command:     "/diet",
				Description: "Set dietary preferences",
			},
			{
				Command:     "/mealplan",
				Description: "Generate a weekly meal plan",
			},
			{
				Command:     "/help",
				Description: "Get detailed instructions",
			},
		},
	})
}

// bot default handler
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	app.handleState(ctx, b, update)

}
