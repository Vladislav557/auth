package repository

import (
	"context"
	"github.com/Vladislav557/auth/internal/models/entity"
	"github.com/Vladislav557/auth/internal/resources/postgres"
	"github.com/google/uuid"
	pg "github.com/lib/pq"
)

const (
	roleUser = "ROLE_USER"
)

type UserRepository struct{}

func (userRepository *UserRepository) ChangeStatusByID(ID int64, status string) error {
	query := `UPDATE users SET status = $1 WHERE id = $2`
	_, err := postgres.DB.Exec(query, status, ID)
	if err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := `
		SELECT id, uuid, full_name, email, password, phone, roles, created_at, updated_at, deleted_at, status
		FROM users WHERE email = $1
	`
	err := postgres.DB.QueryRow(query, email).Scan(
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
		&user.Status)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *UserRepository) GetByUUID(ctx context.Context, UUID string) (*entity.User, error) {
	var user entity.User
	query := `
		SELECT id, uuid, full_name, email, password, phone, roles, created_at, updated_at, deleted_at, status 
		FROM users WHERE uuid = $1
    `
	uuidStr, err := uuid.Parse(UUID)
	if err != nil {
		return &entity.User{}, err
	}
	err = postgres.DB.QueryRowContext(
		ctx,
		query,
		uuidStr,
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
		&user.Status)
	if err != nil {
		return &entity.User{}, err
	}
	return &user, nil
}

func (userRepository *UserRepository) CreateUser(
	ctx context.Context,
	fullName string,
	email string,
	phone string,
	passHash []byte,
	UUID string) (*entity.User, error) {
	var user entity.User
	roles := pg.StringArray{roleUser}
	query := `
		INSERT INTO users (uuid, full_name, email, password, phone, roles) 
		  	VALUES ($1, $2, $3, $4, $5, $6)
		  	RETURNING id, uuid, full_name, email, password, phone, roles, created_at, updated_at, deleted_at
	`
	err := postgres.DB.QueryRowContext(
		ctx,
		query,
		UUID,
		fullName,
		email,
		passHash,
		phone,
		roles,
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
		&user.DeletedAt)
	if err != nil {
		return &entity.User{}, err
	}
	return &user, nil
}
