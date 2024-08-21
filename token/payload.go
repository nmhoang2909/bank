package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("exppired token")
)

type Payload struct {
	// ID       string `json:"id"`
	// Username string `json:"username"`
	jwt.RegisteredClaims
	// IssuedAt  time.Time `json:"issued_at"`
	// ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		// ID:        id.String(),
		// Username:  username,
		// IssuedAt:  time.Now(),
		// ExpiredAt: time.Now().Add(duration),
		jwt.RegisteredClaims{
			Issuer:    username,
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(duration)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ID:        id.String(),
		},
	}, nil
}
