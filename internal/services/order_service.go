package services

import (
	"errors"
	"github.com/hasib-003/orderManagement/internal/models"
	"github.com/hasib-003/orderManagement/internal/repositories"
	"github.com/hasib-003/orderManagement/utils"
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

	if !utils.ValidatePhoneNumber(order.RecipientPhone) {
		return nil, errors.New("invalid phone number")
	}
	order.DeliveryFee = utils.CalculateDeliveryFee(order.RecipientCity, order.ItemWeight)

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
