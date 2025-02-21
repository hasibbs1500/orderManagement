package repositories

import (
	"github.com/hasib-003/orderManagement/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	RegisterUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	LoginUser(email, password string) (*models.User, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
func (repo *UserRepository) RegisterUser(user *models.User) error {
	err := repo.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
