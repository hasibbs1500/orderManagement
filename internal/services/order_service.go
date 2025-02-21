package services

import (
	"github.com/hasib-003/orderManagement/internal/models"
	"github.com/hasib-003/orderManagement/internal/repositories"
)

type OrderServiceInterface interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetAllOrders(limit, page int) ([]*models.Order, int, int, int, int, error)
	CancelOrder(consignmentID string) error
}
type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}
func (service *OrderService) CreateOrder(order *models.Order) (*models.Order, error) {
	if order.RecipientCity == 1 {
		if order.ItemWeight <= 0.5 {
			order.DeliveryFee = 60
		} else if order.ItemWeight > 0.5 && order.ItemWeight <= 1 {
			order.DeliveryFee = 70
		} else {
			order.DeliveryFee = 70 + ((order.ItemWeight - 1) * 15)
		}
	} else {
		order.DeliveryFee = 100 + ((order.ItemWeight - 1) * 15)
	}
	order.CODFee = order.AmountToCollect * 0.01
	err := service.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil

}
func (service *OrderService) GetAllOrders(limit, page int) ([]*models.Order, int, int, int, int, error) {
	orders, total, err := service.repo.GetAllOrders(limit, page)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	lastPage := (total + int64(limit) - 1) / int64(limit)
	return orders, int(total), page, limit, int(lastPage), nil
}
func (service *OrderService) CancelOrder(consignmentID int) error {
	err := service.repo.CancelOrder(consignmentID)
	if err != nil {
		return err
	}
	return nil
}
