package main

import (
	"go-api/api/handlers"
	"go-api/api/repositories"
	"go-api/api/routes"
	"go-api/api/services"
	"go-api/configs"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadEnv()
	app := configs.LoadAppConfig()

	db, err := configs.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	categoriesRepo := repositories.NewCategoriesRepository(db)
	categoriesService := services.NewCategoriesService(categoriesRepo)
	categoriesHandler := handlers.NewCategoriesHandler(categoriesService)

	api := r.Group("/api")
	routes.UserRoutes(api, userHandler)
	routes.CategoriesRoutes(api, categoriesHandler)

	r.Run(":" + app.Port)
}