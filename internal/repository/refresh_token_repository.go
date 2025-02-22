package repository

import (
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/Vladislav557/auth/internal/resources/postgres"
	"github.com/google/uuid"
	"time"
)

type RefreshTokenRepository struct{}

func (refreshTokenRepository *RefreshTokenRepository) GetActiveRefreshToken(user *entity.User, device string) (entity.RefreshToken, error) {
	var refreshToken entity.RefreshToken
	deviceUUID, err := uuid.Parse(device)
	if err != nil {
		return entity.RefreshToken{}, err
	}
	now := time.Now()
	err = postgres.DB.QueryRow(
		"SELECT id, uuid, created_at, expired_at, user_uuid, active, device_uuid FROM refresh_tokens WHERE user_uuid = $1 AND device_uuid = $2 AND active = TRUE AND expired_at > $3",
		user.UUID,
		deviceUUID,
		now,
	).Scan(
		&refreshToken.ID,
		&refreshToken.UUID,
		&refreshToken.CreatedAt,
		&refreshToken.ExpiredAt,
		&refreshToken.UserUUID,
		&refreshToken.Active,
		&refreshToken.DeviceUUID,
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

func (refreshTokenRepository *RefreshTokenRepository) DeactivateRefreshTokenByUserAndDevice(user *entity.User, device string) error {
	deviceUUID, err := uuid.Parse(device)
	if err != nil {
		return err
	}
	_, err = postgres.DB.Exec(
		"UPDATE refresh_tokens SET active = FALSE WHERE user_uuid = $1 AND active = TRUE AND device_uuid = $2",
		user.UUID,
		deviceUUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (refreshTokenRepository *RefreshTokenRepository) CreateRefreshToken(user *entity.User, device string) (entity.RefreshToken, error) {
	token, err := refreshTokenRepository.GetActiveRefreshToken(user, device)
	if err != nil {
		if err = refreshTokenRepository.DeactivateRefreshToken(token); err != nil {
			return entity.RefreshToken{}, err
		}
	}
	if token.ID != 0 {
		return token, nil
	}
	var refreshToken entity.RefreshToken
	expiredAt := time.Now().Add(time.Hour * 24 * 30)
	id := uuid.New().String()
	deviceUUID, err := uuid.Parse(device)
	if err != nil {
		return entity.RefreshToken{}, err
	}
	err = postgres.DB.QueryRow(
		"INSERT INTO refresh_tokens (user_uuid, expired_at, uuid, device_uuid) VALUES ($1, $2, $3, $4) RETURNING id, uuid, created_at, expired_at, uuid, active, device_uuid",
		user.UUID,
		expiredAt,
		id,
		deviceUUID,
	).Scan(
		&refreshToken.ID,
		&refreshToken.UUID,
		&refreshToken.CreatedAt,
		&refreshToken.ExpiredAt,
		&refreshToken.UserUUID,
		&refreshToken.Active,
		&refreshToken.DeviceUUID,
	)
	if err != nil {
		return entity.RefreshToken{}, err
	}
	return refreshToken, nil
}
