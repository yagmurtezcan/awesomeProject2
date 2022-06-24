package server

import (
	"awesomeProject2/internal/handler"
	"awesomeProject2/internal/repository"
	"awesomeProject2/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) StartServer() {
	router := gin.Default()

	userRepository := repository.NewUserRepository(s.db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	loginService := service.NewLoginService(userRepository)
	loginHandler := handler.NewLoginHandler(loginService)

	user := router.Group("/user")
	{
		user.GET("/", userHandler.GetAllUser)
		user.GET("/:id", userHandler.GetUserById)
		user.POST("/", userHandler.CreateUser)
		user.POST("/:id", userHandler.UpdateUser)
		user.DELETE("/:id", userHandler.DeleteUser)
	}

	login := router.Group("/login")
	{
		login.POST("/", loginHandler.Login)
	}

	err := router.Run("localhost:3000")
	if err != nil {
		panic(err)
	}
}
