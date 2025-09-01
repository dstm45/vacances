-- name: CreateUser :one
INSERT INTO users(username, email, password_hash) VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE uuid=$1;

-- name: GetUser :one
SELECT * FROM users WHERE uuid=$1;

-- name: ChangeEmail :exec
UPDATE users SET email=$1 WHERE uuid=$1;

-- name: ChangeUsername :exec
UPDATE users SET username=$1 WHERE uuid=$1;

-- name: ChangePassword :exec
UPDATE users SET password_hash=$1 WHERE uuid=$1;

