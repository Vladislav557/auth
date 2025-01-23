package service

import (
	"github.com/Vladislav557/auth/internal/lib/jwt"
	"github.com/Vladislav557/auth/internal/models/http/request"
	"github.com/Vladislav557/auth/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	rr repository.RegistrationRepository
}

func (rs *RegistrationService) Register(req request.RegistrationRequest) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	id := uuid.New().String()
	if err != nil {
		return "", err
	}
	u, err := rs.rr.Registration(req.FullName, req.Email, req.Phone, passHash, id)
	if err != nil {
		return "", err
	}
	token, err := jwt.NewToken(u)
	if err != nil {
		return "", err
	}
	return token, nil
}
