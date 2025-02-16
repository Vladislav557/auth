package jwt

import (
	"errors"
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func NewAccessToken(user *entity.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["UUID"] = user.UUID
	claims["roles"] = user.Roles
	claims["status"] = user.Status
	claims["exp"] = time.Now().Add(time.Minute * 30)

	return toString(token)
}

func NewRefreshToken(refresh *entity.RefreshToken) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["UUID"] = refresh.UUID
	claims["exp"] = refresh.ExpiredAt
	claims["createdAt"] = refresh.CreatedAt
	claims["active"] = refresh.Active

	return toString(token)
}

func toString(token *jwt.Token) (string, error) {
	path := os.Getenv("JWT_PRIVATE_KEY")
	secret, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("secret key not found")
	}

	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
