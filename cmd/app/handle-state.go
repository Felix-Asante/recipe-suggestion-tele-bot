package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/botStates"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"gorm.io/gorm"
)

func (app *application) handleState(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println(update.CallbackQuery.Data)
	if update.Message == nil {
		return
	}
	chatId := update.Message.Chat.ID
	chatState, err := app.repositories.BotState.FindByChatId(chatId)

	if nil != err && !errors.Is(err, gorm.ErrRecordNotFound) {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    chatId,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}

	switch chatState.State {
	case botStates.WaitingForDietPreference:
		app.handleDietPreference(ctx, b, update)
	case botStates.WaitingForPantry:
		fmt.Println("Waiting for pantry")
	case botStates.WaitingForMealPlan:
		fmt.Println("Waiting for meal plan")
	case botStates.WaitingForRecipeSearch:
		fmt.Println("Waiting for recipe search")
	default:
		app.startHandler(ctx, b, update)
	}
}
