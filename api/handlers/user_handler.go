package handlers

import (
	"go-api/api/components/response"
	"go-api/api/entities"
	"go-api/api/requests"
	"go-api/api/services"
	"net/http"
	"strconv"
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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users, total, err := h.service.GetPaginated(c, page, limit)
	if err != nil {
		code, res := response.Error(
			http.StatusInternalServerError,
			"Failed to get users",
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

	code, res := response.SuccessWithMeta("Success", users, meta)
	c.JSON(code, res)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func (h *UserHandler) Show(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user, err := h.service.GetByID(c, id)
	if err != nil {
		code, res := response.Error(500, "Failed", err.Error())
		c.JSON(code, res)
		return
	}

	if user == nil {
		code, res := response.Error(404, "User not found", nil)
		c.JSON(code, res)
		return
	}

	code, res := response.Success("Success", user)
	c.JSON(code, res)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if err := h.service.Delete(c, id); err != nil {
		code, res := response.Error(500, "Failed delete", err.Error())
		c.JSON(code, res)
		return
	}

	code, res := response.Success("User deleted", nil)
	c.JSON(code, res)
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		code, res := response.Error(
			http.StatusBadRequest,
			"Invalid user ID",
			nil,
		)
		c.JSON(code, res)
		return
	}

	var req requests.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		code, res := response.ValidationError(err.Error())
		c.JSON(code, res)
		return
	}

	user := entities.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.service.Update(c, id, user); err != nil {
		code, res := response.Error(
			http.StatusInternalServerError,
			"Failed to update user",
			err.Error(),
		)
		c.JSON(code, res)
		return
	}

	code, res := response.Success("User updated", nil)
	c.JSON(code, res)
}
