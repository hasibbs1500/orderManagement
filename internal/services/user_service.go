package services

import (
	"errors"
	"github.com/hasib-003/orderManagement/internal/models"
	"github.com/hasib-003/orderManagement/internal/repositories"
	"github.com/hasib-003/orderManagement/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserServiceInterface interface {
	RegisterUser(user *models.User) error
	LoginUser(email, password string) (string, error)
}
type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}
func (s *UserService) RegisterUser(user *models.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	err = s.repo.RegisterUser(user)
	if err != nil {
		return errors.New("cannot register user")
	}
	return nil
}
func (s *UserService) LoginUser(email, password string) (models.LoginResponse, error) {
	if email == "" || password == "" {
		return models.LoginResponse{}, errors.New("email and password are required")
	}
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return models.LoginResponse{}, errors.New("cannot login user")
	}
	if user == nil {
		return models.LoginResponse{}, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.LoginResponse{}, errors.New("invalid password")
	}
	token, err := utils.GenerateToken(user.Email, time.Hour*5)
	if err != nil {
		return models.LoginResponse{}, errors.New("cannot generate token")
	}
	return token, nil
}
