package tododomain

type Repository interface {
	FetchTodo(todoID int) (*Todo, error)
	FetchAllTodo() ([]*Todo, error)
	CreateTodo(todo *Todo) (int, error)
	UpdateTodo(todo *Todo) error
	DeleteTodo(todoID int) error
}
