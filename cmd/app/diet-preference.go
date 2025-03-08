package main

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/botStates"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/dto"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/repositories"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (app *application) dietHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	createStateDto := dto.CreateBotStateDto{
		ChatId: update.Message.Chat.ID,
		State:  botStates.WaitingForDietPreference,
	}
	if _, err := app.repositories.BotState.Upsert(createStateDto); nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}

	messageParams := &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      messages.SetDietaryPreference,
		ParseMode: models.ParseModeMarkdownV1,
	}
	if _, err := b.SendMessage(ctx, messageParams); nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}
}

func (app *application) handleDietPreference(ctx context.Context, b *bot.Bot, update *models.Update) {
	preferences := update.Message.Text
	userId := update.Message.From.ID
	preferenceList, invalidPreferences := extractPreferences(preferences)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      messages.SomethingWentWrong,
				ParseMode: models.ParseModeMarkdownV1,
			})
		}
	}()

	// save preferences
	if len(preferenceList) > 0 {
		dietPreference := make([]*repositories.DietPreference, 0)
		for _, preference := range preferenceList {
			dietPreference = append(dietPreference, &repositories.DietPreference{
				UserId:     userId,
				Preference: preference,
			})
		}

		if err := app.repositories.DietPreference.Create(dietPreference); nil != err {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      messages.SomethingWentWrong,
				ParseMode: models.ParseModeMarkdownV1,
			})
			return
		}
	}

	if err := app.repositories.BotState.RemoveByChatId(update.Message.Chat.ID); nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}

	// send success message with invalid preferences
	message := getMessage(invalidPreferences, preferenceList)
	messageParams := &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      message,
		ParseMode: models.ParseModeMarkdownV1,
	}
	if _, err := b.SendMessage(ctx, messageParams); nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}

}

func getMessage(invalidPreferences []string, preferenceList []string) string {
	message := fmt.Sprintf(messages.InvalidDietaryPreference, strings.Join(invalidPreferences, ","))
	if len(preferenceList) > 0 && len(invalidPreferences) > 0 {
		message = fmt.Sprintf(messages.DietaryPreferenceSavedWithoutInvalid, strings.Join(invalidPreferences, ","))
	}
	if len(preferenceList) > 0 && len(invalidPreferences) == 0 {
		message = messages.DietaryPreferencesSaved
	}
	return message
}

func extractPreferences(preferences string) ([]string, []string) {
	re := regexp.MustCompile(`^[a-zA-Z\s]+(,[a-zA-Z\s]+)*$`)
	invalidPreferences := []string{}
	validPreferences := []string{}

	preferenceList := strings.Split(preferences, ",")
	for _, pref := range preferenceList {
		pref = strings.TrimSpace(pref)
		if !re.MatchString(pref) || pref == "" {
			invalidPreferences = append(invalidPreferences, pref)
		} else {
			validPreferences = append(validPreferences, pref)
		}
	}
	return validPreferences, invalidPreferences
}
