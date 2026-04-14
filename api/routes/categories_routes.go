package routes

import (
	"go-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(r *gin.RouterGroup, h *handlers.CategoriesHandler) {
	r.GET("/categories", h.GetAll)
	r.GET("/categories/:id", h.ShowCategories)
}