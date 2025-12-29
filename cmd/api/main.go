package main

import (
	"log"

	"go-api/internal/app/user"
	"go-api/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := user.NewDB(
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// Feature-based route
	user.RegisterRoutes(r, db)

	log.Println("server running on :" + cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
