-- name: CreateEntry :execlastid
INSERT INTO entries
(account_id, amount) VALUES (?, ?);

-- name: GetEntryById :one
select * from entries where id = ?;

-- name: DeleteEntryByAccountId :exec
DELETE FROM entries WHERE account_id = ?;
