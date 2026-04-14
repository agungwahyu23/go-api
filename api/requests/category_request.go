package requests

type CreateCategoriesRequest struct {
	Name        string    `json:"name" binding:"required"`
}