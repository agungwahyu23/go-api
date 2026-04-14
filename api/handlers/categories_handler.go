package handlers

import (
	"go-api/api/components/response"
	"go-api/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoriesHandler struct {
	service services.CategoriesService
}

func NewCategoriesHandler(service services.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{service}
}

func (h *CategoriesHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	categories, total, err := h.service.GetAllCategories(c, page, limit)
	if err != nil {
		code, res := response.Error(
			http.StatusInternalServerError,
			"Failed to get categories",
			err.Error(),
		)
		c.JSON(code, res)
		return
	}

	meta := response.PaginationMeta{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	code, res := response.SuccessWithMeta("Success", categories, meta)
	c.JSON(code, res)
}

func (h *CategoriesHandler) ShowCategories(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	categories, err := h.service.GetByIDCategories(c, id)
	if err != nil {
		code, res := response.Error(500, "Failed", err.Error())
		c.JSON(code, res)
		return
	}

	if categories == nil {
		code, res := response.Error(404, "Data not found", nil)
		c.JSON(code, res)
		return
	}

	code, res := response.Success("Success", categories)
	c.JSON(code, res)
}