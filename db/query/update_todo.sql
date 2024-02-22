-- name: UpdateTodo :one
UPDATE todos
SET title = $2, done = $3
WHERE id = $1
RETURNING id, title, done;

