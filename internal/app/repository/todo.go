package repository

import (
	"errors"

	"github.com/massanaRoger/todo-go-htmx/internal/model"
)

// TodoRepository interface for the todo storage.
type TodoRepository interface {
	Add(todo model.Todo) (model.Todo, error)
	Get() ([]model.Todo, error)
	GetById(id int) (model.Todo, error)
}

// InMemoryTodoRepository in-memory repo implementation.
type InMemoryTodoRepository struct {
	todos []model.Todo
}

func NewInMemoryTodoRepository() *InMemoryTodoRepository {
	return &InMemoryTodoRepository{}
}

func (r *InMemoryTodoRepository) Add(todo model.Todo) (model.Todo, error) {
	for i, t := range r.todos {
		if t.ID == todo.ID {
			r.todos[i] = todo
			return todo, nil
		}
	}

	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *InMemoryTodoRepository) Get() ([]model.Todo, error) {
	return r.todos, nil
}

func (r *InMemoryTodoRepository) GetById(id int) (model.Todo, error) {
	for _, item := range r.todos {
		if item.ID == id {
			return item, nil
		}
	}
	return r.todos[0], errors.New("ID not found")
}
