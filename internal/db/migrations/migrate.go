package main

import (
	"log"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db"
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/repositories"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	models := []interface{}{
		&repositories.User{},
		&repositories.BotState{},
		&repositories.DietPreference{},
		&repositories.SavedRecipe{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal(err)
	}
}
