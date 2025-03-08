package db

import (
	"fmt"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", env.GetString("DB_HOST", "localhost"), env.GetString("DB_USER", "postgres"), env.GetString("DB_PASSWORD", "pgdev"), env.GetString("DB_NAME", "recipe_pal_bot_db"), env.GetString("DB_PORT", "5432"), env.GetString("SSL_MODE", "disable"))
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
