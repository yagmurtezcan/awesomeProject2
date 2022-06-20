package handler

import (
	"awesomeProject2/internal/dto"
	"awesomeProject2/internal/model"
	"awesomeProject2/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) GetAllUser(c *gin.Context) {
	user := u.userService.GetAllUser()

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

func (u *UserHandler) DeleteUser(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Param("id"))

}
