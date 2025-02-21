package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hasib-003/orderManagement/internal/models"
	"os"
	"time"
)

var (
	revokedTokens = make(map[string]bool)
	jwtKey        = []byte(os.Getenv("JWT_KEY"))
)

type Claims struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(email string, expiry time.Duration) (models.LoginResponse, error) {
	expirationTime := time.Now().Add(expiry).Unix()
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return models.LoginResponse{}, err
	}
	return models.LoginResponse{
		TokenType:   "Bearer",
		Expiresin:   expirationTime,
		AccessToken: signedToken,
	}, nil
}
func RevokeToken(token string) {
	revokedTokens[token] = true
}
