package jwt

import (
	"errors"
	"github.com/Vladislav557/auth/internal/models/domain"
	"github.com/Vladislav557/auth/internal/models/entity"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

func GetAccessToken(user *entity.User) (string, error) {
	token := jwtlib.New(jwtlib.SigningMethodHS256)
	claims := token.Claims.(jwtlib.MapClaims)
	claims["sub"] = user.UUID
	claims["name"] = user.FullName
	claims["roles"] = user.Roles
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["iss"] = os.Getenv("SERVICE_UUID")

	return toString(token)
}

func GetRefreshToken(refresh *entity.RefreshToken) (string, error) {
	return refresh.UUID, nil
}

func ParseToken(tokenStr string) (domain.Claims, error) {
	var user domain.Claims
	path := os.Getenv("JWT_PRIVATE_KEY")
	secret, err := os.ReadFile(path)
	if err != nil {
		return user, errors.New("secret key not found")
	}
	tokenWithoutBearer := strings.TrimPrefix(tokenStr, "Bearer ")
	token, err := jwtlib.ParseWithClaims(tokenWithoutBearer, &domain.Claims{}, func(token *jwtlib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return user, err
	}
	if claims, ok := token.Claims.(*domain.Claims); ok && token.Valid {
		return *claims, nil
	}
	return user, errors.New("token not valid")
}

func toString(token *jwtlib.Token) (string, error) {
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
