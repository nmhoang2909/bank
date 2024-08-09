-- name: CreateTransfer :execlastid
insert into transfers
(from_account_id, to_account_id, amount) values (?, ?, ?);

-- name: GetTransferById :one
select * from transfers where id = ?;
