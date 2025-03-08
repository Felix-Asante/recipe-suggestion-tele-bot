package main

import (
	"context"
	"fmt"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (app *application) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userId := update.Message.From.ID
	if update.Message.From.IsBot {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "I'm a bot, you can't start me",
			ParseMode: models.ParseModeMarkdown,
		})
	}

	// create / update user
	fmt.Println(userId)
	// send welcome message
	sendWelcomeMessage(ctx, b, update)

}

func sendWelcomeMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := fmt.Sprintf(messages.WelcomeText, update.Message.From.FirstName)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      message,
		ParseMode: models.ParseModeMarkdownV1,
	})
}
