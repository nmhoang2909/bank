package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndCheckPw(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		hashedPw, err := HashPassword("password")
		assert.NoError(t, err)
		assert.NotZero(t, hashedPw)

		ok, err := IsCorrectPassword([]byte(hashedPw), []byte("password"))
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("incorrect", func(t *testing.T) {
		hashedPw, err := HashPassword("password")
		assert.NoError(t, err)
		assert.NotZero(t, hashedPw)

		ok, err := IsCorrectPassword([]byte(hashedPw), []byte("password 1"))
		assert.NotNil(t, err)
		assert.False(t, ok)
	})
}
