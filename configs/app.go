package configs

import "os"

type AppConfig struct {
	Env  string
	Port string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		Env:  os.Getenv("APP_ENV"),
		Port: os.Getenv("APP_PORT"),
	}
}