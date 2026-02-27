package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret")

type AuthUsecase struct{}

func (a *AuthUsecase) Login(username, password string) (string, error) {
	if username != "admin" || password != "password" {
		return "", errors.New("invalid credentials")
	}
	return a.GenerateToken(username)
}

func (a *AuthUsecase) GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
