package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/massanaRoger/todo-go-htmx/db"
)

type TodoRepository interface {
	Add(todo db.Todo) (db.Todo, error)
	Get() ([]db.Todo, error)
	GetById(id int32) (db.Todo, error)
	RemoveTodo(id int32) error
	EditTodo(todo db.Todo) (db.Todo, error)
}

type PostgresTodoRepository struct {
	queries *db.Queries
	todos   []db.Todo
}

func NewPostgresTodoRepository(conn *pgx.Conn) *PostgresTodoRepository {
	queries := db.New(conn)
	return &PostgresTodoRepository{queries: queries}
}

func (r *PostgresTodoRepository) Add(todo db.Todo) (db.Todo, error) {
	insertedTodo, err := r.queries.CreateTodo(context.Background(), db.CreateTodoParams{
		Title: todo.Title,
		Done:  todo.Done,
	})
	if err != nil {
		return db.Todo{}, err
	}
	return insertedTodo, nil
}

func (r *PostgresTodoRepository) Get() ([]db.Todo, error) {
	todos, err := r.queries.ListTodos(context.Background())
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *PostgresTodoRepository) GetById(id int32) (db.Todo, error) {
	todo, err := r.queries.GetTodoByID(context.Background(), id)
	if err != nil {
		return db.Todo{}, err
	}

	return todo, nil
}

func (r *PostgresTodoRepository) RemoveTodo(id int32) error {
	return r.queries.DeleteTodo(context.Background(), id)
}

func (r *PostgresTodoRepository) EditTodo(todo db.Todo) (db.Todo, error) {
	return r.queries.UpdateTodo(context.Background(), db.UpdateTodoParams{
		Title: todo.Title,
		Done:  todo.Done,
		ID:    todo.ID,
	})
}
