package tododomain

type Repository interface {
	FetchTodo(todoID int) (*Todo, error)
	FetchAllTodo() ([]*Todo, error)
}
