package main

import (
	"context"
	"fmt"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/dto"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (app *application) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.From.IsBot {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "I'm a bot, you can't start me",
			ParseMode: models.ParseModeMarkdown,
		})
		return
	}

	// create / update user
	userDto := dto.CreateUserDto{
		FirstName: update.Message.From.FirstName,
		LastName:  update.Message.From.LastName,
		UserId:    update.Message.From.ID,
	}
	if err := app.repositories.User.Upsert(userDto); err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "Something went wrong, please try again later",
			ParseMode: models.ParseModeMarkdown,
		})
		return
	}
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
