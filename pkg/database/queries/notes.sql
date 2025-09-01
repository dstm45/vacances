-- name: CreateNote :one
INSERT INTO notes(title, author_id) VALUES($1, $2)
RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes WHERE uuid=$1;

-- name: GetNote :one
SELECT * FROM notes WHERE uuid=$1;


