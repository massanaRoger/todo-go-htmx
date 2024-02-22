package model

type TodoForm struct {
	ID    int    `form:"id"`
	Title string `form:"title"`
	Done  bool   `form:"done"`
}

type CheckTodo struct {
	ID int32 `form:"id"`
}

type EditTodo struct {
	ID        int32  `form:"id"`
	PrevValue string `form:"prevValue"`
	NewValue  string `form:"newValue"`
}

type RemoveTodo struct {
	ID int32 `form:"id"`
}
