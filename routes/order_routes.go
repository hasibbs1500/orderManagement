package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/orderManagement/config"
	"github.com/hasib-003/orderManagement/internal/handlers"
	"github.com/hasib-003/orderManagement/internal/repositories"
	"github.com/hasib-003/orderManagement/internal/services"
	"github.com/hasib-003/orderManagement/middleware"
)

func RegisterOrderRoutes(router *gin.Engine) {
	db := config.GetDB()
	orderRepository := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepository)
	orderHandler := handlers.NewOrderHandler(orderService)
	api := router.Group("/api/v1")
	api.POST("/createOrder", middleware.TokenValidationMiddleware(), orderHandler.CreateOrder)
	api.GET("/getAllOrders", middleware.TokenValidationMiddleware(), orderHandler.GetAllOrders)
	api.PUT("/orders/:id/cancel", middleware.TokenValidationMiddleware(), orderHandler.CancelOrder)
}
