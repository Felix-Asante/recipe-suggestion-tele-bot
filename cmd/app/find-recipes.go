package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"strings"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/ai"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/botStates"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/dto"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/repositories"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/messages"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/utils"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/utils/structs"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/slider"
	"github.com/samber/lo"
)

func (app *application) findRecipesHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	defer utils.RecoverFromPanic()
	userId := update.Message.From.ID
	dietPreferences, err := app.repositories.DietPreference.FindByUserId(userId)

	if nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}
	if len(dietPreferences) == 0 {
		app.setWaitingForDietPreference(ctx, b, update)
		return
	}

	createBotStateDto := dto.CreateBotStateDto{
		ChatId: update.Message.Chat.ID,
		State:  botStates.WaitingForRecipeSearch,
	}
	if err := app.repositories.BotState.Upsert(createBotStateDto); nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}

	messageParams := &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      messages.FindRecipes,
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

func (app *application) handleRecipeSearch(ctx context.Context, b *bot.Bot, update *models.Update) {
	defer utils.RecoverFromPanic()
	ingredients := update.Message.Text
	userId := update.Message.From.ID

	var wg sync.WaitGroup

	slidesChannel := make(chan []slider.Slide)

	dietPreferences, err := app.repositories.DietPreference.FindByUserId(userId)
	if nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}

	preferences := lo.Map(dietPreferences, func(dietPreference repositories.DietPreference, _ int) string {
		return dietPreference.Preference
	})

	dietPreference := strings.Join(preferences, ", ")
	aiPrompt := strings.ReplaceAll(ai.FIND_RECIPE_PROMPT, "{userIngredients}", ingredients)
	aiPrompt = strings.ReplaceAll(aiPrompt, "{userDietaryPreference}", dietPreference)

	aiResponse, err := ai.GenerateRecipes(aiPrompt)
	if nil != err {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    update.Message.Chat.ID,
			Text:      messages.SomethingWentWrong,
			ParseMode: models.ParseModeMarkdownV1,
		})
		return
	}

	go func() {
		for _, recipe := range aiResponse {
			wg.Add(1)
			go func(r structs.Recipe) {
				defer wg.Done()
				generateSlides(r, slidesChannel)
			}(recipe)
		}
		wg.Wait()
		close(slidesChannel)
	}()

	slides := []slider.Slide{}

	for s := range slidesChannel {
		slides = append(slides, s...)
	}

	opts := []slider.Option{
		slider.OnSelect("❤️ Save", false, sliderOnSelect),
	}

	sl := slider.New(b, slides, opts...)

	_, err = sl.Show(ctx, b, update.Message.Chat.ID)
	if nil != err {
		fmt.Println("error showing slider")
		fmt.Println(err)
		// b.SendMessage(ctx, &bot.SendMessageParams{
		// 	ChatID:    update.Message.Chat.ID,
		// 	Text:      messages.SomethingWentWrong,
		// 	ParseMode: models.ParseModeMarkdownV1,
		// })
		return
	}
}

func generateSlides(recipe structs.Recipe, slidesChannel chan []slider.Slide) {
	photo, err := ai.GeneratePhotos(recipe.Title)
	recipeContent := formatRecipeContent(recipe)
	if nil != err {
		fmt.Println("error generating photos")
		fmt.Println(err)
		slidesChannel <- []slider.Slide{}
		return
	}
	slide := slider.Slide{
		Text:  recipeContent,
		Photo: photo,
	}
	slidesChannel <- []slider.Slide{slide}
}

func formatRecipeContent(recipe structs.Recipe) string {
	content := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s",
		recipe.Title,
		recipe.Ingredients,
		recipe.Instructions,
		recipe.DietaryCompliance)

	specialChars := []string{
		"-", "_", "*", "[", "]", "(", ")", "~", "`", ">",
		"#", "+", "=", "|", "{", "}", ".", "!",
	}

	for _, char := range specialChars {
		content = strings.ReplaceAll(content, char, "\\"+char)
	}

	return content
}

func sliderOnSelect(ctx context.Context, b *bot.Bot, message models.MaybeInaccessibleMessage, item int) {
	userId := message.Message.Chat.ID
	fmt.Println("Selected item:", userId)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: message.Message.Chat.ID,
		Text:   "Select " + strconv.Itoa(item),
	})
}

func sliderOnCancel(ctx context.Context, b *bot.Bot, message models.MaybeInaccessibleMessage) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: message.Message.Chat.ID,
		Text:   "Cancel",
	})
}
