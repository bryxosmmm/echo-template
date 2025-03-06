-- name: CreateClient :one
insert into client (email, password_hash, name)
values ($1, $2, $3)
returning *;

-- name: GetClient :one
select * from client
where email = $1;


