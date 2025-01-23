package domain

import "time"

type JwtUser struct {
	ID        int64
	UUID      string
	FullName  string
	Email     string
	Phone     string
	Roles     []string
	CreatedAt time.Time
}
