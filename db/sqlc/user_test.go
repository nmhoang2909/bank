package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/nmhoang2909/bank/util"
	"github.com/stretchr/testify/assert"
)

func randomCreateUser(t *testing.T) (username, hashesPw string) {
	username = util.RandomString(6)
	pw := util.RandomString(5)
	hashedPw, _ := util.HashPassword(pw)
	_, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Username:     username,
		FullName:     util.RandomString(10),
		Email:        fmt.Sprintf("%s@email.com", util.RandomString(5)),
		HashPassword: hashedPw,
	})
	assert.NoError(t, err)

	return
}
