package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/orderManagement/internal/models"
	"github.com/hasib-003/orderManagement/internal/services"
	"github.com/hasib-003/orderManagement/utils"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"ID":    user.ID,
		"email": user.Email,
	})
}
func (h *UserHandler) LoginUser(c *gin.Context) {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	token, err := h.service.LoginUser(userInput.Email, userInput.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The user credentials were incorrect.",
			"type": "error",
			"code": 400})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
func (h *UserHandler) LogoutUser(c *gin.Context) {
	var request struct {
		AccessToken string `json:"access_token"`
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	utils.RevokeToken(request.AccessToken)
	c.JSON(http.StatusOK, gin.H{"Message": "logged out successfully"})
}
