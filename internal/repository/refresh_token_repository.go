package repository

import (
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/Vladislav557/auth/internal/resources/postgres"
	"github.com/google/uuid"
	"time"
)

type RefreshTokenRepository struct{}

func (refreshTokenRepository *RefreshTokenRepository) getActiveRefreshToken(user *entity.User) (entity.RefreshToken, error) {
	var refreshToken entity.RefreshToken
	err := postgres.DB.QueryRow(
		"SELECT id, uuid, created_at, expired_at, user_id, active FROM refresh_tokens WHERE user_id = ($1) AND active = TRUE",
		user.ID,
	).Scan(
		&refreshToken.ID,
		&refreshToken.UUID,
		&refreshToken.CreatedAt,
		&refreshToken.ExpiredAt,
		&refreshToken.UserID,
		&refreshToken.Active,
	)
	if err != nil {
		return refreshToken, err
	}
	return refreshToken, nil
}

func (refreshTokenRepository *RefreshTokenRepository) deactivateAllRefreshTokens(user *entity.User) error {
	_, err := postgres.DB.Exec(
		"UPDATE refresh_tokens SET active = FALSE WHERE user_id = ($1) AND active = TRUE",
		user.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (refreshTokenRepository *RefreshTokenRepository) CreateRefreshToken(user *entity.User) (entity.RefreshToken, error) {
	token, err := refreshTokenRepository.getActiveRefreshToken(user)
	if err == nil {
		return token, nil
	}
	err = refreshTokenRepository.deactivateAllRefreshTokens(user)
	if err != nil {
		return entity.RefreshToken{}, err
	}
	var refreshToken entity.RefreshToken
	expiredAt := time.Now().Add(time.Hour * 48)
	id := uuid.New().String()
	err = postgres.DB.QueryRow(
		"INSERT INTO refresh_tokens (user_id, expired_at, uuid) VALUES ($1, $2, $3) RETURNING id",
		user.ID,
		expiredAt,
		id,
	).Scan(&refreshToken.ID)
	if err != nil {
		return entity.RefreshToken{}, err
	}
	err = postgres.DB.QueryRow("SELECT id, uuid, created_at, expired_at, user_id, active FROM refresh_tokens WHERE id = $1",
		refreshToken.ID).Scan(
		&refreshToken.ID,
		&refreshToken.UUID,
		&refreshToken.CreatedAt,
		&refreshToken.ExpiredAt,
		&refreshToken.UserID,
		&refreshToken.Active,
	)
	if err != nil {
		return entity.RefreshToken{}, err
	}
	return refreshToken, nil
}
