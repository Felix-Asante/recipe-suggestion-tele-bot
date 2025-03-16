package main

import (
	"context"
	"fmt"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/slider"
)

func (app *application) savedRecipesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	userId := update.Message.From.ID
	savedRecipes, err := app.repositories.SavedRecipe.FindByUserId(userId)
	if nil != err {
		fmt.Println("error finding saved recipes", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}
	if len(savedRecipes) == 0 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.NoSavedRecipes,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}

	slides := make([]slider.Slide, len(savedRecipes))
	for i, recipe := range savedRecipes {
		slides[i] = slider.Slide{
			Photo:    recipe.Photo,
			Text:     formatRecipeContent(recipe.Caption),
			IsUpload: false,
		}
	}

	sl := slider.New(b, slides, slider.OnSelect("❌ Delete", false, app.sliderOnDeleteSavedRecipe))
	_, err = sl.Show(ctx, b, update.Message.Chat.ID)
	if nil != err {
		fmt.Println("error showing slider", err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}
}

func (app *application) sliderOnDeleteSavedRecipe(ctx context.Context, b *bot.Bot, message models.MaybeInaccessibleMessage, item int) {
	photoID := message.Message.Photo[0].FileID
	if err := app.repositories.SavedRecipe.RemoveByPhoto(photoID); nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    message.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: message.Message.Chat.ID,
		Text:   "Deleted",
	})

	app.savedRecipesHandler(ctx, b, &models.Update{
		Message: &models.Message{
			Chat: message.Message.Chat,
			From: &models.User{
				ID: message.Message.Chat.ID,
			},
		},
	})
}
