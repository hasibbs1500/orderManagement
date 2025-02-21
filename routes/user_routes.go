package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/orderManagement/config"
	"github.com/hasib-003/orderManagement/internal/handlers"
	"github.com/hasib-003/orderManagement/internal/repositories"
	"github.com/hasib-003/orderManagement/internal/services"
)

func RegisterUserRoutes(router *gin.Engine) {
	db := config.GetDB()
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/logout", userHandler.LogoutUser)

}
