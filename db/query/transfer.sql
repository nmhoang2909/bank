-- name: CreateTransfer :execlastid
insert into transfers
(from_account_id, to_account_id, amount) values (?, ?, ?);

-- name: GetTransferById :one
select * from transfers where id = ?;

-- name: DeleteTransferByFromAccontId :exec
DELETE FROM transfers WHERE from_account_id = ?;

-- name: DeleteTransferByToAccontId :exec
DELETE FROM transfers WHERE to_account_id = ?;

