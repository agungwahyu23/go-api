package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	AppName string
	AppPort string

	DBDriver	string
	DBHost		string
	DBPort		string
	DBName		string
	DBUser		string
	DBPassword	string
}

func Load() *Config  {
	_ = godotenv.Load()

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),

		DBDriver: os.Getenv("DB_DRIVER"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}

	if cfg.AppPort == "" {
		log.Fatal("APP_PORT is required")
	}

	return cfg
}

func NewDB(user, password, host, port, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbname,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, db.Ping()
}