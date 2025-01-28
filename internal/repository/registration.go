package repository

import (
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/Vladislav557/auth/internal/resources/postgres"
	"github.com/google/uuid"
	pg "github.com/lib/pq"
	"time"
)

type RegistrationRepository struct{}

func (rr *RegistrationRepository) Registration(
	fullName string,
	email string,
	phone string,
	passHash []byte,
	uuid string) (entity.User, error) {
	var user entity.User
	roles := pg.StringArray{"ROLE_USER"}
	err := postgres.DB.QueryRow(
		"INSERT INTO users (uuid, full_name, email, password, phone, roles) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		uuid,
		fullName,
		email,
		passHash,
		phone,
		roles,
	).Scan(&user.ID)
	if err != nil {
		return entity.User{}, err
	}

	// Выполняем SELECT для получения полной записи пользователя
	err = postgres.DB.QueryRow(
		"SELECT id, uuid, full_name, email, password, phone, roles, created_at, updated_at, deleted_at FROM users WHERE id = $1",
		user.ID,
	).Scan(
		&user.ID,
		&user.UUID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Roles,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (rr *RegistrationRepository) CreateRefreshToken(user entity.User) (entity.RefreshToken, error) {
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
