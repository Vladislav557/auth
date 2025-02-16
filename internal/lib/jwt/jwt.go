package jwt

import (
	"errors"
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/golang-jwt/jwt/v5"
	pg "github.com/lib/pq"
	"os"
	"time"
)

type UserClaims struct {
	UUID      string         `json:"UUID"`
	Roles     pg.StringArray `json:"roles"`
	Status    string         `json:"status"`
	ExpiredAt time.Time      `json:"exp"`
	jwt.RegisteredClaims
}

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

func Parse(tokenStr string) (UserClaims, error) {
	var user UserClaims
	path := os.Getenv("JWT_PRIVATE_KEY")
	secret, err := os.ReadFile(path)
	if err != nil {
		return user, errors.New("secret key not found")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return user, err
	} else if claims, ok := token.Claims.(*UserClaims); ok {
		return *claims, nil
	} else {
		return user, nil
	}
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
