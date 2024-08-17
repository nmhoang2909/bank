package db

import (
	"context"
	"testing"

	"github.com/nmhoang2909/bank/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	username := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    username,
		Balance:  200,
		Currency: "USD",
	}
	id, err := testQueries.CreateAccount(context.Background(), arg)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}

func createRandomAccount(t *testing.T) (accountId int64) {
	username := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    username,
		Balance:  int32(util.RandomNumber(500, 10000)),
		Currency: util.RandomCurrency(),
	}
	id, _ := testQueries.CreateAccount(context.Background(), arg)
	return id
}
