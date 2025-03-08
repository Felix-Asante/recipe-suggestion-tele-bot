package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key string, fallback string) string {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to load .env file: %v\n", err)
		return fallback
	}
	return os.Getenv(key)
}

func GetInt(key string, fallback int) int {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to load .env file: %v\n", err)
		return fallback
	}
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return fallback
	}
	return value
}
