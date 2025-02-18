package service

import (
	"context"
	"errors"
	"github.com/Vladislav557/auth/internal/lib/jwt"
	"github.com/Vladislav557/auth/internal/models/http/request"
	"github.com/Vladislav557/auth/internal/models/http/response"
	"github.com/Vladislav557/auth/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthorizationService struct {
	userRepository         repository.UserRepository
	refreshTokenRepository repository.RefreshTokenRepository
	notifier               Notifier
}

//func (authorizationService *AuthorizationService) Logout(claims domain.Claims) error {
//
//}

func (authorizationService *AuthorizationService) LoginByEmail(email string, password string) (response.TokenResponse, error) {
	user, err := authorizationService.userRepository.GetByEmail(email)
	if err != nil {
		return response.TokenResponse{}, err
	}
	if user.Status == "new" {
		return response.TokenResponse{}, errors.New("user not confirmed by email")
	}
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return response.TokenResponse{}, errors.New("password incorrect")
	}
	refreshToken, err := authorizationService.refreshTokenRepository.CreateRefreshToken(user)
	if err != nil {
		return response.TokenResponse{}, err
	}
	accessTokenStr, err := jwt.GetAccessToken(user)
	refreshTokenStr, err := jwt.GetRefreshToken(&refreshToken)
	if err != nil {
		return response.TokenResponse{}, err
	}
	return response.TokenResponse{AccessToken: accessTokenStr, RefreshToken: refreshTokenStr}, nil
}

func (authorizationService *AuthorizationService) Register(ctx context.Context, req request.SingUpRequest) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	id := uuid.New().String()
	if err != nil {
		return err
	}
	user, err := authorizationService.userRepository.CreateUser(ctx, req.FullName, req.Email, req.Phone, passHash, id)
	if err = authorizationService.notifier.AcceptRegistration(user); err != nil {
		return err
	}
	return nil
}

func (authorizationService *AuthorizationService) Confirm(ctx context.Context, UUID string) error {
	user, err := authorizationService.userRepository.GetByUUID(ctx, UUID)
	if err != nil {
		return err
	}
	err = authorizationService.userRepository.ChangeStatusByID(user.ID, "confirmed")
	return nil
}
