package service

import (
	"github.com/massanaRoger/todo-go-htmx/internal/app/repository"
	"github.com/massanaRoger/todo-go-htmx/internal/model"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) AddTodo(todo model.Todo) (model.Todo, error) {
	// Business logic (validation, etc.) goes here
	return s.repo.Add(todo)
}

func (s *TodoService) GetTodos() ([]model.Todo, error) {
	return s.repo.Get()
}

func (s *TodoService) GetTodo(id int) (model.Todo, error) {
	return s.repo.GetById(id)
}

func (s *TodoService) RemoveTodo(id int) error {
	return s.repo.RemoveTodo(id)
}
