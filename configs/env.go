package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Failed load .env")
		}
	}
}
