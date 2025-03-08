package main

import (
	"context"
	"fmt"
	"net/http"

	"os"
	"os/signal"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/env"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

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

	storage := db.NewStorage()

	app := &application{bot: b, storage: storage}

	app.run()

	go func() {
		http.ListenAndServe(env.GetString("PORT", ":2000"), b.WebhookHandler())
	}()

	b.StartWebhook(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// b.SendMessage(ctx, &bot.SendMessageParams{
	// 	ChatID:    update.Message.Chat.ID,
	// 	Text:      "Hello, World!",
	// 	ParseMode: models.ParseModeMarkdown,
	// })
}
