package entity

import (
	pg "github.com/lib/pq"
	"time"
)

type User struct {
	ID        string         `json:"id"`
	UUID      string         `json:"uuid"`
	FullName  string         `json:"full_name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Password  []byte         `json:"password"`
	Roles     pg.StringArray `json:"roles"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt interface{}    `json:"deleted_at"`
}
