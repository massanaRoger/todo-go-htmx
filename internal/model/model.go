package model

type Todo struct {
	ID    int    `form:"id"`
	Title string `form:"title"`
	Done  bool   `form:"done"`
}

type CheckTodo struct {
	ID int `form:"id"`
}

type EditTodo struct {
	ID    int    `form:"id"`
	Value string `form:"value"`
}
