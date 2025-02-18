package domain

import (
	jwtlib "github.com/golang-jwt/jwt/v5"
	pg "github.com/lib/pq"
	"time"
)

type Claims struct {
	Roles                   pg.StringArray `json:"roles"`
	Exp                     time.Duration  `json:"exp"`
	Iat                     time.Duration  `json:"iat"`
	Iss                     string         `json:"iss"`
	Name                    string         `json:"name"`
	Nbf                     time.Duration  `json:"nbf"`
	Sub                     string         `json:"sub"`
	jwtlib.RegisteredClaims `json:"-"`
}
