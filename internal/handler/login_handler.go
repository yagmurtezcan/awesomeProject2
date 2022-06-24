package handler

import (
	"awesomeProject2/internal/model"
	"awesomeProject2/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	loginService *service.LoginService
}

func NewLoginHandler(loginService *service.LoginService) *LoginHandler {
	return &LoginHandler{loginService: loginService}
}

func (l *LoginHandler) Login(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		return
	}

	loginUser, errs := l.loginService.Login(user)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": loginUser})
}
