-- name: CreateAccount :execlastid
INSERT INTO accounts (
  owner, balance, currency
) VALUES (
  ?, ?, ?
);

-- name: GetAccountById :one
SELECT * FROM accounts WHERE id = ? LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM accounts LIMIT ? OFFSET ?;

-- name: DeleteAccountById :exec
DELETE FROM accounts WHERE id = ?;

-- name: UpdateAccountById :exec
UPDATE accounts
  SET balance = ?
  WHERE id = ?;

-- name: UpdateAccountBalance :exec
UPDATE accounts
  SET balance = balance + sqlc.arg(amount)
  WHERE id = sqlc.arg(id);
