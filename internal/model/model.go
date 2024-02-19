package model

type Todo struct {
	ID    int    `form:"id"`
	Title string `form:"title"`
	Done  bool   `form:"done"`
}

type CheckTodo struct {
	Checked bool `form:"checked"`
	ID      int  `form:"id"`
}
