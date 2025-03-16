package main

import (
	"context"
	"fmt"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (app *application) helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.From.IsBot {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      "I'm a bot, you can't start me",
			ParseMode: models.ParseModeMarkdown,
		})
		return
	}

	messageParams := &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      messages.HelpText,
		ParseMode: models.ParseModeHTML,
	}

	if _, err := b.SendMessage(ctx, messageParams); nil != err {
		fmt.Println("error sending help message", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}
}
