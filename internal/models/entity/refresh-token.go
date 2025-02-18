package entity

import (
	"time"
)

type RefreshToken struct {
	ID        int64     `json:"id" db:"id"`
	UserUUID  string    `json:"user_uuid" db:"user_uuid"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiredAt time.Time `json:"expired_at" db:"expired_at"`
	UUID      string    `json:"uuid" db:"uuid"`
	Active    bool      `json:"active" db:"active"`
}
