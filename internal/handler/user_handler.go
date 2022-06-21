package handler

import (
	"awesomeProject2/internal/dto"
	"awesomeProject2/internal/model"
	"awesomeProject2/internal/service"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var userDTO dto.UserDTO
	c.BindJSON(&userDTO)

	model := model.User{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
	}

	createdUser, err := u.userService.CreateUser(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

func (u *UserHandler) GetUserById(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}

	user, err := u.userService.GetUserById(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) DeleteUser(c *gin.Context) {
	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}

	err := u.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Operation completed"})
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	var userDTO dto.UserDTO
	c.BindJSON(&userDTO)

	id, error := strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}

	userDTO.ID = id

	updateUser, err := u.userService.UpdateUser(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, &updateUser)
}
