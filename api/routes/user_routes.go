package routes

import (
	"go-api/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, h *handlers.UserHandler) {
	r.POST("/users", h.Create)
	r.GET("/users", h.GetAll)
}