package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/botStates"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"gorm.io/gorm"
)

func (app *application) handleState(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.CallbackQuery != nil {
		app.handleCallbackQuery(ctx, b, update)
		return
	}
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

func (app *application) handleCallbackQuery(ctx context.Context, b *bot.Bot, update *models.Update) {
	query := update.CallbackQuery.Data
	chat := update.CallbackQuery.Message.Message.Chat

	parts := strings.Split(query, "_")

	if len(parts) < 3 {
		b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       true,
			Text:            messages.InvalidCallback,
		})
		return
	}

	action := parts[0] + "_" + parts[1]

	switch action {
	case "delete_diet-preference":
		app.handleDeleteDietPreference(ctx, b, update)
	case "add_diet-preference":
		newUpdate := &models.Update{
			Message: &models.Message{
				Chat: chat,
			},
		}
		b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
		})
		app.setWaitingForDietPreference(ctx, b, newUpdate)
	}
}
