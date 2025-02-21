package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}
type LoginResponse struct {
	TokenType   string `json:"token_type"`
	Expiresin   int64  `json:"expires_in"`
	AccessToken string `json:"access_token"`
}
