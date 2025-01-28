package service

import (
	"github.com/Vladislav557/auth/internal/lib/jwt"
	"github.com/Vladislav557/auth/internal/models/http/request"
	"github.com/Vladislav557/auth/internal/models/http/response"
	"github.com/Vladislav557/auth/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	rr repository.RegistrationRepository
}

func (rs *RegistrationService) Register(req request.RegistrationRequest) (response.RegistrationResponse, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	id := uuid.New().String()
	if err != nil {
		return response.RegistrationResponse{}, err
	}
	u, err := rs.rr.Registration(req.FullName, req.Email, req.Phone, passHash, id)
	if err != nil {
		return response.RegistrationResponse{}, err
	}
	rt, err := rs.rr.CreateRefreshToken(u)
	if err != nil {
		return response.RegistrationResponse{}, err
	}
	accessToken, err := jwt.NewAccessToken(u)
	if err != nil {
		return response.RegistrationResponse{}, err
	}
	refreshToken, err := jwt.NewRefreshToken(rt)
	if err != nil {
		return response.RegistrationResponse{}, err
	}
	return response.RegistrationResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
