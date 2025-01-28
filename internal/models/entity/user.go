package entity

import (
	pg "github.com/lib/pq"
	"time"
)

type User struct {
	ID        string
	UUID      string
	FullName  string
	Email     string
	Phone     string
	Password  []byte
	Roles     pg.StringArray
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt interface{}
}
