package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to load .env file: %v\n", err)
		return ""
	}
	return os.Getenv(key)
}
