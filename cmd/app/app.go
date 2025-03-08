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

	// app.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHandler)
}

// func startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
// 	userId := update.Message.From.ID
// 	if update.Message.From.IsBot {
// 		b.SendMessage(ctx, &bot.SendMessageParams{
// 			ChatID:    update.Message.Chat.ID,
// 			Text:      "I'm a bot, you can't start me",
// 			ParseMode: models.ParseModeMarkdown,
// 		})
// 	}
// 	fmt.Println(userId)
// 	b.SendMessage(ctx, &bot.SendMessageParams{
// 		ChatID:    update.Message.Chat.ID,
// 		Text:      "Hello, *" + bot.EscapeMarkdown(update.Message.From.FirstName) + "*",
// 		ParseMode: models.ParseModeMarkdown,
// 	})
// }
