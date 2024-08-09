package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
)

type Store struct {
	db *sql.DB
	*Queries
	mu sync.Mutex
}

func NewStore(db *sql.DB) *Store {
	return &Store{db, New(db), sync.Mutex{}}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := s.Queries.WithTx(tx)
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rollback err: %v, exec err: %v", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}

type (
	TransferTxParams struct {
		FromAccountID int64 `db:"from_account_id"`
		ToAccountID   int64 `db:"to_account_id"`
		Amount        int64 `db:"amount"`
	}
	TransferTxResult struct {
		Transfer    Transfer `db:"transfer"`
		FromAccount Account  `db:"from_account"`
		ToAccount   Account  `db:"to_account"`
		FromEntry   Entry    `db:"from_entry"`
		ToEntry     Entry    `db:"to_entry"`
	}
)

func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		// create transfer for history tracing
		transferId, err := s.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        int32(arg.Amount),
		})
		if err != nil {
			return fmt.Errorf("create transfer err: %v", err)
		}
		result.Transfer, err = s.GetTransferById(ctx, transferId)
		if err != nil {
			return err
		}

		// create from entry
		fromEntryId, err := s.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -int32(arg.Amount),
		})
		if err != nil {
			return fmt.Errorf("create entry err: %v", err)
		}
		result.FromEntry, err = s.GetEntryById(ctx, fromEntryId)
		if err != nil {
			return err
		}

		// create to entry
		toEntryId, err := s.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    int32(arg.Amount),
		})
		if err != nil {
			return fmt.Errorf("create entry err: %v", err)
		}
		result.ToEntry, err = s.GetEntryById(ctx, toEntryId)
		if err != nil {
			return err
		}

		// update balance accounts
		// fromAccount, err := s.Queries.GetAccountByIdForUpdate(ctx, arg.FromAccountID)
		// if err != nil {
		// 	return err
		// }
		// fromAccToBeUpdate := UpdateAccountByIdParams{
		// 	Balance: fromAccount.Balance - int32(arg.Amount),
		// 	ID:      arg.FromAccountID,
		// }
		// if err = s.Queries.UpdateAccountById(ctx, fromAccToBeUpdate); err != nil {
		// 	return err
		// }
		if err = s.Queries.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
			Amount: -int32(arg.Amount),
			ID:     arg.FromAccountID,
		}); err != nil {
			return err
		}

		fromAccount, err := s.Queries.GetAccountById(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}
		result.FromAccount = fromAccount

		// toAccount, err := s.Queries.GetAccountByIdForUpdate(ctx, arg.ToAccountID)
		// if err != nil {
		// 	return err
		// }
		// toAccToBeUpdate := UpdateAccountByIdParams{
		// 	Balance: toAccount.Balance + int32(arg.Amount),
		// 	ID:      arg.ToAccountID,
		// }
		// if err = s.Queries.UpdateAccountById(ctx, toAccToBeUpdate); err != nil {
		// 	return err
		// }

		if err = s.Queries.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
			Amount: int32(arg.Amount),
			ID:     arg.ToAccountID,
		}); err != nil {
			return err
		}
		toAccount, err := s.Queries.GetAccountById(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}
		result.ToAccount = toAccount

		return nil
	})

	return result, err
}
