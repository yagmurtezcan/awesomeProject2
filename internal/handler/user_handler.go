package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *UserHandler) GetAllUser(c *gin.Context) {
	user := u.userService.GetAllUser()

	c.JSON(http.StatusOK, gin.H{"user": user})
}
