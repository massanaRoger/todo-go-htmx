-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;

