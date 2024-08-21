package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateJWTToken(t *testing.T) {
	jwt, err := NewJWTMaker("abcxyz")
	assert.NoError(t, err)

	token, err := jwt.CreateToken("hoang", time.Minute)
	assert.NoError(t, err)
	assert.NotZero(t, token)
}
