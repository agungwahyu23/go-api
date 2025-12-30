package handlers

import (
	"go-api/api/components/response"
	"go-api/api/entities"
	"go-api/api/requests"
	"go-api/api/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req requests.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		code, res := response.ValidationError(err.Error())
		c.JSON(code, res)
		return
	}

	password, _ := HashPassword(req.Password)
	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
        // Handle error jika format tanggal salah
        return 
    }

	user := entities.User{
		Name:  req.Name,
		Email: req.Email,
		Username: req.Username,
		Password: password,
		Address: req.Address,
		Phone: req.Phone,
		DateOfBirth: dob,
		Gender: req.Gender,
		IsActive: req.IsActive,
	}

	if err := h.service.Create(c, user); err != nil {
		code, res := response.Error(
			http.StatusInternalServerError,
			"Failed to create user",
			err.Error(),
		)
		c.JSON(code, res)
		return
	}

	code, res := response.Created("User created", user)
	c.JSON(code, res)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll(c)
	if err != nil {
		code, res := response.Error(
			http.StatusInternalServerError,
			"Failed to get users",
			err.Error(),
		)
		c.JSON(code, res)
		return
	}

	code, res := response.Success("Success", users)
	c.JSON(code, res)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}