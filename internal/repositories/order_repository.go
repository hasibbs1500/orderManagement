package repositories

import (
	"github.com/hasib-003/orderManagement/internal/models"
	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	CreateOrder(order *models.Order) error
	GetAllOrders(limit int, page int) ([]models.Order, int64, error)
	CancelOrder(consignmentId string) error
}
type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}
func (repo *OrderRepository) CreateOrder(order *models.Order) error {
	err := repo.db.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}
func (repo *OrderRepository) GetAllOrders(limit int, page int) ([]*models.Order, int64, error) {
	var orders []*models.Order
	var total int64
	query := repo.db.Model(&models.Order{}).Where("order_status != ?", "canceled")
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit > 0 && page > 0 {
		offset := (page - 1) * limit
		query = query.Limit(limit).Offset(offset)
	}
	err = query.Find(&orders).Error
	if err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}
func (repo *OrderRepository) CancelOrder(consignmentId int) error {
	order := repo.db.Model(&models.Order{}).Where("id=?", consignmentId).Update("order_status", "canceled")
	if order.Error != nil {
		return order.Error
	}
	if order.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
