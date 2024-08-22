package token

import (
	"testing"
	"time"

	"github.com/nmhoang2909/bank/util"
	"github.com/stretchr/testify/assert"
)

func TestPasetoMaker(t *testing.T) {
	symmetricKey := "d7d8271949084258b9c29f19272c802b"
	t.Run("valid", func(t *testing.T) {
		username := util.RandomString(5)
		maker, err := NewPasetoMaker(symmetricKey)
		assert.NoError(t, err)

		token, err := maker.CreateToken(username, time.Minute)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		payload, err := maker.VerifyToken(token)
		assert.NoError(t, err)
		assert.NotNil(t, payload)
		assert.Equal(t, username, payload.Issuer)
		assert.WithinDuration(t, time.Now().Add(time.Minute), payload.ExpiresAt.Time, time.Second)
	})
}
