-- name: CreateUser :execlastid
insert into users (username, full_name, email, hash_password) values (?, ?, ?, ?);

-- name: GetUser :one
select * from users where username = ?;
