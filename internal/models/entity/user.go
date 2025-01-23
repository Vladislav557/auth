package entity

import "time"

type User struct {
	ID        string
	UUID      string
	FullName  string
	Email     string
	Phone     string
	Password  []byte
	Roles     []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt interface{}
}
