package db

import (
	"context"
	"testing"

	"github.com/nmhoang2909/bank/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "BM",
		Balance:  200,
		Currency: "USD",
	}
	id, err := testQueries.CreateAccount(context.Background(), arg)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}

func createRandomAccount() (accountId int64) {
	arg := CreateAccountParams{
		Owner:    util.RandomString(6),
		Balance:  int32(util.RandomNumber(500, 10000)),
		Currency: util.RandomCurrency(),
	}
	id, _ := testQueries.CreateAccount(context.Background(), arg)
	return id
}
