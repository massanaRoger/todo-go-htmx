package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/massanaRoger/todo-go-htmx/db"
	"github.com/massanaRoger/todo-go-htmx/internal/app/service"
	"github.com/massanaRoger/todo-go-htmx/internal/app/templates"
	"github.com/massanaRoger/todo-go-htmx/internal/app/util"
	"github.com/massanaRoger/todo-go-htmx/internal/model"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) AddTodo(c echo.Context) error {
	var newTodo model.TodoForm
	if err := c.Bind(&newTodo); err != nil {
		return err
	}

	todo := db.Todo{
		Title: newTodo.Title,
		Done:  false,
	}
	addedTodo, err := h.service.AddTodo(todo)
	if err != nil {
		return err
	}
	return util.Render(c, 200, templates.Todo(addedTodo))
}

func (h *TodoHandler) AllTodos(c echo.Context) error {
	todos, err := h.service.GetTodos()
	if err != nil {
		return err
	}
	return util.Render(c, 200, templates.AllTodos(todos))
}

func (h *TodoHandler) CheckTodo(c echo.Context) error {
	var todoToCheck model.CheckTodo
	if err := c.Bind(&todoToCheck); err != nil {
		return err
	}

	todo, err := h.service.GetTodo(todoToCheck.ID)
	if err != nil {
		return err
	}
	editedTodo, err := h.service.EditTodo(db.Todo{
		Done:  !todo.Done,
		Title: todo.Title,
		ID:    todoToCheck.ID,
	})
	if err != nil {
		return err
	}
	return util.Render(c, 200, templates.CheckTodo(editedTodo))
}

func (h *TodoHandler) StartEditTodo(c echo.Context) error {
	qp := c.QueryParam("id")
	id_int, err := strconv.Atoi(qp)
	id := int32(id_int)
	if err != nil {
		return err
	}
	todo, err := h.service.GetTodo(id)
	if err != nil {
		return err
	}
	return util.Render(c, 200, templates.StartEditTodo(todo))
}

func (h *TodoHandler) EditTodo(c echo.Context) error {
	var editTodo model.EditTodo
	if err := c.Bind(&editTodo); err != nil {
		return err
	}

	trigger, err := h.whatTrigger(c)
	if err != nil {
		return err
	}

	var todo db.Todo
	todo.ID = editTodo.ID

	if trigger == "blur" {
		todo.Title = editTodo.PrevValue
	} else if trigger == "keyup" {
		todo.Title = editTodo.NewValue
	} else {
		return errors.New("Unknown trigger")
	}

	editedTodo, err := h.service.EditTodo(todo)
	if err != nil {
		return err
	}

	return util.Render(c, 200, templates.EditTodo(editedTodo))
}

func (h *TodoHandler) RemoveTodo(c echo.Context) error {
	var removeTodo model.RemoveTodo
	if err := c.Bind(&removeTodo); err != nil {
		return err
	}

	err := h.service.RemoveTodo(removeTodo.ID)

	if err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "")
}

func (h *TodoHandler) whatTrigger(c echo.Context) (string, error) {
	headers := c.Request().Header
	header := headers.Get("Triggering-Event")
	var eventData map[string]interface{}
	if err := json.Unmarshal([]byte(header), &eventData); err != nil {
		return "", err
	}

	if htmxInternalData, ok := eventData["htmx-internal-data"].(map[string]interface{}); ok {
		if triggerSpec, ok := htmxInternalData["triggerSpec"].(map[string]interface{}); ok {
			if trigger, ok := triggerSpec["trigger"].(string); ok {
				return trigger, nil
			}
		}
	}

	return "", errors.New("Trigger not found")
}
