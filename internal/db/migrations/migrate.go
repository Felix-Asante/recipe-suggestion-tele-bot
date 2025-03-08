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

	if err := db.AutoMigrate(&repositories.User{}, &repositories.BotState{}); err != nil {
		log.Fatal(err)
	}
}
