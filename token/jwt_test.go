package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nmhoang2909/bank/util"
	"github.com/stretchr/testify/assert"
)

const secretKey = "secret"

func TestJWTMaker(t *testing.T) {
	var username = util.RandomString(5)
	t.Run("valid", func(t *testing.T) {
		maker, err := NewJWTMaker(secretKey)
		assert.NoError(t, err)

		duration := time.Minute
		issuedAt := time.Now()
		expiredAt := issuedAt.Add(duration)

		token, err := maker.CreateToken(username, time.Minute)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		payload, err := maker.VerifyToken(token)
		assert.NoError(t, err)
		assert.NotNil(t, payload)

		assert.Equal(t, payload.Issuer, username)
		assert.NotZero(t, payload.ID)
		assert.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
		assert.WithinDuration(t, expiredAt, payload.ExpiresAt.Time, time.Second)
	})

	t.Run("expired", func(t *testing.T) {
		maker, err := NewJWTMaker(secretKey)
		assert.NoError(t, err)

		token, err := maker.CreateToken(username, -time.Minute)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		payload, err := maker.VerifyToken(token)
		assert.Error(t, err)
		assert.Nil(t, payload)
		assert.ErrorIs(t, err, ErrExpiredToken)
	})

	t.Run("invalid", func(t *testing.T) {
		payload, err := NewPayload(username, time.Minute)
		assert.NoError(t, err)

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
		token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
		assert.NoError(t, err)

		maker, err := NewJWTMaker(secretKey)
		assert.NoError(t, err)
		p, err := maker.VerifyToken(token)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidToken)
		assert.Nil(t, p)
	})

}
