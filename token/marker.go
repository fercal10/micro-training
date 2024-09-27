package token

import (
	"time"
)

const token = "12345678901234567890123456789012"

func NewMarker() Maker {

	tokenMaker, _ := NewJWTMaker(token)
	return tokenMaker
}

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
