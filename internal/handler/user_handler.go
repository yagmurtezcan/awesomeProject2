package handler

import (
	"awesomeProject2/internal/dto"
	"awesomeProject2/internal/model"
	"awesomeProject2/internal/service"
	"awesomeProject2/internal/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) GetAllUser(c *gin.Context) {
	user, err := u.userService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var userDTO dto.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		return
	}

	validationError := validator.ValidateRequest(userDTO)
	if validationError != nil {
		c.JSON(http.StatusBadRequest, validationError)
		return
	}

	user := model.User{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
	}

	createdUser, err := u.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

func (u *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Operation completed"})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var userDTO dto.UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		return
	}

	validationError := validator.ValidateRequest(userDTO)
	if validationError != nil {
		c.JSON(http.StatusBadRequest, validationError)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO.ID = id

	updateUser, err := u.userService.UpdateUser(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &updateUser)
}
