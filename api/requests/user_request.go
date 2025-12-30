package requests

type CreateUserRequest struct {
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	DateOfBirth string `json:"date_of_birth"`
	Gender int `json:"gender"`
	IsActive int `json:"is_active"`
}