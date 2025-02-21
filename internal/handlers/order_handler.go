package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hasib-003/orderManagement/internal/models"
	"github.com/hasib-003/orderManagement/internal/services"
	"log"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{orderService}

}
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Please fix the given errors", "type": "error", "code": 422, "errors": err.Error()})
		return
	}
	createOrder, err := h.orderService.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "failed to create order"})
		return
	}
	response := gin.H{"Message": "order created successfully", "type": "success", "code": 200, "data": gin.H{
		"consignment_id": createOrder.ID,
		"merchant_id":    createOrder.MerchantOrderID,
		"order_status":   createOrder.OrderStatus,
		"delevery_fee":   createOrder.DeliveryFee,
	}}
	c.JSON(http.StatusOK, response)
}
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	Page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if limit < 1 || Page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid pagination Parameters"})
		return
	}
	orders, total, currentPage, perPage, lastPage, err := h.orderService.GetAllOrders(limit, Page)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "failed to get orders"})
		return
	}
	response := gin.H{
		"message": "Orders successfully fetched.",
		"type":    "success",
		"code":    200,
		"data": gin.H{
			"data":          orders,
			"total":         total,
			"current_page":  currentPage,
			"per_page":      perPage,
			"total_in_page": len(orders),
			"last_page":     lastPage,
		},
	}
	c.JSON(http.StatusOK, response)
}
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.orderService.CancelOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "failed to cancel order"})
	}
	c.JSON(http.StatusOK, gin.H{"Message": "order successfully canceled"})

}
