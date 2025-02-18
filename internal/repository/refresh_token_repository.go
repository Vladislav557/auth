package repository

import (
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/Vladislav557/auth/internal/resources/postgres"
	"github.com/google/uuid"
	"time"
)

type RefreshTokenRepository struct{}

func (refreshTokenRepository *RefreshTokenRepository) GetActiveRefreshToken(user *entity.User) (entity.RefreshToken, error) {
	var refreshToken entity.RefreshToken
	err := postgres.DB.QueryRow(
		"SELECT id, uuid, created_at, expired_at, uuid, active FROM refresh_tokens WHERE user_id = ($1) AND active = TRUE",
		user.ID,
	).Scan(
		&refreshToken.ID,
		&refreshToken.UUID,
		&refreshToken.CreatedAt,
		&refreshToken.ExpiredAt,
		&refreshToken.UserUUID,
		&refreshToken.Active,
	)
	if err != nil {
		return refreshToken, err
	}
	return refreshToken, nil
}

func (refreshTokenRepository *RefreshTokenRepository) DeactivateRefreshToken(refreshToken entity.RefreshToken) error {
	_, err := postgres.DB.Exec(
		"UPDATE refresh_tokens SET active = FALSE WHERE id = ($1)",
		refreshToken.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (refreshTokenRepository *RefreshTokenRepository) DeactivateAllRefreshTokens(user *entity.User) error {
	_, err := postgres.DB.Exec(
		"UPDATE refresh_tokens SET active = FALSE WHERE user_uuid = ($1) AND active = TRUE",
		user.UUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (refreshTokenRepository *RefreshTokenRepository) CreateRefreshToken(user *entity.User) (entity.RefreshToken, error) {
	if token, err := refreshTokenRepository.GetActiveRefreshToken(user); err != nil {
		if err = refreshTokenRepository.DeactivateRefreshToken(token); err != nil {
			return entity.RefreshToken{}, err
		}
	}
	var refreshToken entity.RefreshToken
	expiredAt := time.Now().Add(time.Hour * 24 * 30)
	id := uuid.New().String()
	err := postgres.DB.QueryRow(
		"INSERT INTO refresh_tokens (user_uuid, expired_at, uuid) VALUES ($1, $2, $3) RETURNING id, uuid, created_at, expired_at, uuid, active",
		user.UUID,
		expiredAt,
		id,
	).Scan(
		&refreshToken.ID,
		&refreshToken.UUID,
		&refreshToken.CreatedAt,
		&refreshToken.ExpiredAt,
		&refreshToken.UserUUID,
		&refreshToken.Active,
	)
	if err != nil {
		return entity.RefreshToken{}, err
	}
	return refreshToken, nil
}
