package entity

import (
	"time"
)

type RefreshToken struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiredAt time.Time `json:"expired_at" db:"expired_at"`
	UUID      string    `json:"uuid" db:"uuid"`
	Active    bool      `json:"active" db:"active"`
}
