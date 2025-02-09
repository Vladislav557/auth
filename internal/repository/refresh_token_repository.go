package repository

import (
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/Vladislav557/auth/internal/resources/postgres"
	"github.com/google/uuid"
	"time"
)

type RefreshTokenRepository struct{}

func (refreshTokenRepository *RefreshTokenRepository) CreateRefreshToken(user *entity.User) (entity.RefreshToken, error) {
	var refreshToken entity.RefreshToken
	expiredAt := time.Now().Add(time.Hour * 48)
	id := uuid.New().String()
	err := postgres.DB.QueryRow(
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
