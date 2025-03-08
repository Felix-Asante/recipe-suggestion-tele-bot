package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"os"
	"os/signal"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/repositories"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/env"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var app *application

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

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

// bot default handler
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	app.handleState(ctx, b, update)

}
