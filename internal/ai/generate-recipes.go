package ai

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/env"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/utils/structs"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateRecipes(prompt string) ([]structs.Recipe, error) {
	ctx := context.Background()
	apiKey := env.GetString("AI_API_KEY", "")

	if apiKey == "" {
		return nil, errors.New("AI_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	recipes := extractRecipes(resp)

	return recipes, nil
}

func extractRecipes(resp *genai.GenerateContentResponse) []structs.Recipe {
	recipes := make([]structs.Recipe, 0)

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				text, ok := part.(genai.Text)
				if !ok {
					fmt.Println("Part is not of type genai.Text")
					continue
				}
				cleanedResponse := cleanResponse(string(text))
				var recipeList []structs.Recipe
				if err := json.Unmarshal([]byte(cleanedResponse), &recipeList); err != nil {
					fmt.Println("Error unmarshaling recipe:", err)
					continue
				}

				recipes = append(recipes, recipeList...)
			}
		}
	}

	return recipes
}

func cleanResponse(response string) string {
	cleaned := strings.ReplaceAll(response, "```json", "")
	cleaned = strings.ReplaceAll(cleaned, "```", "")
	cleaned = strings.TrimSpace(cleaned)
	return cleaned
}
