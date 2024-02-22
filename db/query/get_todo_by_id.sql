-- name: GetTodoByID :one
SELECT id, title, done FROM todos
WHERE id = $1;

