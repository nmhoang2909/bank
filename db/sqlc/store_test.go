package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	accId1 := createRandomAccount()
	accId2 := createRandomAccount()
	acc1, _ := testQueries.GetAccountById(context.Background(), accId1)
	acc2, _ := testQueries.GetAccountById(context.Background(), accId2)

	errs := make(chan error)
	results := make(chan TransferTxResult)
	n := 5
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: accId1,
				ToAccountID:   accId2,
				Amount:        10,
			})
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		assert.NoError(t, err)

		result := <-results
		assert.NotEmpty(t, result)

		// check transfers
		transfer := result.Transfer
		assert.NotEmpty(t, transfer, "transfer empty")
		assert.Equal(t, accId1, transfer.FromAccountID)
		assert.Equal(t, accId2, transfer.ToAccountID)
		assert.Equal(t, int32(10), transfer.Amount)

		// check entries
		fromEntry := result.FromEntry
		assert.NotEmpty(t, fromEntry, "fromEntry empty")
		assert.Equal(t, accId1, fromEntry.AccountID)
		assert.Equal(t, int32(-10), fromEntry.Amount)

		toEntry := result.ToEntry
		assert.NotEmpty(t, toEntry, "toEntry empty")
		assert.Equal(t, accId2, toEntry.AccountID)
		assert.Equal(t, int32(10), toEntry.Amount)

		// check accounts
		fromAcc := result.FromAccount
		assert.NotEmpty(t, fromAcc)
		assert.Equal(t, accId1, fromAcc.ID)
	}
	acc1Updated, _ := testQueries.GetAccountById(context.Background(), accId1)
	assert.Equal(t, acc1.Balance-10*int32(n), acc1Updated.Balance)
	acc2Updated, _ := testQueries.GetAccountById(context.Background(), accId2)
	assert.Equal(t, acc2.Balance+10*int32(n), acc2Updated.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)
	accId1 := createRandomAccount()
	accId2 := createRandomAccount()
	acc1, _ := testQueries.GetAccountById(context.Background(), accId1)
	acc2, _ := testQueries.GetAccountById(context.Background(), accId2)

	errs := make(chan error)
	n := 10
	for i := 0; i < n; i++ {
		fromAccId := accId1
		toAccId := accId2

		if i%2 == 1 {
			fromAccId = accId2
			toAccId = accId1
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccId,
				ToAccountID:   toAccId,
				Amount:        10,
			})
			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		assert.NoError(t, err)
	}
	acc1Updated, _ := testQueries.GetAccountById(context.Background(), accId1)
	assert.Equal(t, acc1.Balance, acc1Updated.Balance)
	acc2Updated, _ := testQueries.GetAccountById(context.Background(), accId2)
	assert.Equal(t, acc2.Balance, acc2Updated.Balance)
}
