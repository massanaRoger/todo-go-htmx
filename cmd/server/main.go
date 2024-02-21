package main

import (
	"github.com/labstack/echo/v4"
	"github.com/massanaRoger/todo-go-htmx/internal/app/handler"
	"github.com/massanaRoger/todo-go-htmx/internal/app/repository"
	"github.com/massanaRoger/todo-go-htmx/internal/app/service"
)

func main() {
	e := echo.New()
	e.Static("/", "public")
	todoRepo := repository.NewInMemoryTodoRepository()
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)
	e.GET("/start-edit-todo", todoHandler.StartEditTodo)

	e.POST("/add-todo", todoHandler.AddTodo)
	e.POST("/check-todo", todoHandler.CheckTodo)
	e.POST("/remove-todo", todoHandler.RemoveTodo)

	e.PUT("/edit-todo", todoHandler.EditTodo)

	e.Logger.Fatal(e.Start(":1323"))
}
