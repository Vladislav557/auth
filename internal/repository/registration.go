package repository

import (
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/Vladislav557/auth/internal/resources/postgres"
)

type RegistrationRepository struct{}

func (rr *RegistrationRepository) Registration(
	fullName string,
	email string,
	phone string,
	passHash []byte,
	uuid string) (entity.User, error) {
	var user entity.User
	err := postgres.DB.QueryRow(
		"INSERT INTO users (uuid, full_name, email, password, phone, roles) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		uuid,
		fullName,
		email,
		passHash,
		phone,
		`{"ROLE_USER"}`,
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
