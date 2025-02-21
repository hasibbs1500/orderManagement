package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/orderManagement/config"
	"github.com/hasib-003/orderManagement/internal/models"
	"github.com/hasib-003/orderManagement/routes"
	"log"
)

func main() {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.User{}, &models.Order{})
	if err != nil {
		log.Println("auto migrate err:", err)
	}
	router := gin.Default()
	routes.RegisterUserRoutes(router)
	routes.RegisterOrderRoutes(router)
	router.Run(":8080")

}
