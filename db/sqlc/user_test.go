package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/nmhoang2909/bank/util"
	"github.com/stretchr/testify/assert"
)

func createRandomUser(t *testing.T) (username string) {
	username = util.RandomString(6)
	_, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Username:     username,
		FullName:     util.RandomString(10),
		Email:        fmt.Sprintf("%s@email.com", util.RandomString(5)),
		HashPassword: "secret",
	})
	assert.NoError(t, err)

	return username
}
