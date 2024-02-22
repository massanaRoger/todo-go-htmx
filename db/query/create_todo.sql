-- name: CreateTodo :one
INSERT INTO todos (title, done)
VALUES ($1, $2)
RETURNING id, title, done;

