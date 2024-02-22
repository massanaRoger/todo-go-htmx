package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/massanaRoger/todo-go-htmx/internal/app/handler"
	"github.com/massanaRoger/todo-go-htmx/internal/app/repository"
	"github.com/massanaRoger/todo-go-htmx/internal/app/service"
)

func main() {
	e := echo.New()
	e.Static("/", "public")

	conn, err := pgx.Connect(context.Background(), "postgres://root:root@localhost:5432/database")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	todoRepo := repository.NewPostgresTodoRepository(conn)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)
	e.GET("/start-edit-todo", todoHandler.StartEditTodo)
	e.GET("/all-todos", todoHandler.AllTodos)

	e.POST("/add-todo", todoHandler.AddTodo)
	e.POST("/check-todo", todoHandler.CheckTodo)
	e.POST("/remove-todo", todoHandler.RemoveTodo)

	e.PUT("/edit-todo", todoHandler.EditTodo)

	e.Logger.Fatal(e.Start(":1323"))
}
