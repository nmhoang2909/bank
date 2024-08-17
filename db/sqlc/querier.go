// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (int64, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (int64, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	DeleteAccountById(ctx context.Context, id int64) error
	DeleteEntryByAccountId(ctx context.Context, accountID int64) error
	DeleteTransferByFromAccontId(ctx context.Context, fromAccountID int64) error
	DeleteTransferByToAccontId(ctx context.Context, toAccountID int64) error
	GetAccountById(ctx context.Context, id int64) (Account, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]Account, error)
	GetEntryById(ctx context.Context, id int64) (Entry, error)
	GetTransferById(ctx context.Context, id int64) (Transfer, error)
	GetUser(ctx context.Context, username string) (User, error)
	UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) error
	UpdateAccountById(ctx context.Context, arg UpdateAccountByIdParams) error
}

var _ Querier = (*Queries)(nil)
