package service

import (
	"github.com/massanaRoger/todo-go-htmx/db"
	"github.com/massanaRoger/todo-go-htmx/internal/app/repository"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) AddTodo(todo db.Todo) (db.Todo, error) {
	return s.repo.Add(todo)
}

func (s *TodoService) GetTodos() ([]db.Todo, error) {
	return s.repo.Get()
}

func (s *TodoService) GetTodo(id int32) (db.Todo, error) {
	return s.repo.GetById(id)
}

func (s *TodoService) RemoveTodo(id int32) error {
	return s.repo.RemoveTodo(id)
}

func (s *TodoService) EditTodo(todo db.Todo) (db.Todo, error) {
	return s.repo.EditTodo(todo)
}
