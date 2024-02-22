// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: get_todo_by_id.sql

package db

import (
	"context"
)

const getTodoByID = `-- name: GetTodoByID :one
SELECT id, title, done FROM todos
WHERE id = $1
`

func (q *Queries) GetTodoByID(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodoByID, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.Done)
	return i, err
}