package usecase

import (
	"github.com/onituka/todo-sample/domain/tododomain"
	"github.com/onituka/todo-sample/usecase/output"
)

type TodoUsecase interface {
	FetchTodo(todoID int) (*output.Todo, error)
}

type todoUsecase struct {
	todoRepository tododomain.Repository
}

func NewTodoUsecase(todoRepository tododomain.Repository) *todoUsecase {
	return &todoUsecase{
		todoRepository: todoRepository,
	}
}

func (u *todoUsecase) FetchTodo(todoID int) (*output.Todo, error) {
	todo, err := u.todoRepository.FetchTodo(todoID)
	if err != nil {
		return nil, err
	}

	return &output.Todo{
		ID:                 todo.ID(),
		Title:              todo.Title(),
		Memo:               todo.Memo(),
		ImplementationDate: output.OutDate{Time: todo.ImplementationDate()},
		DueDate:            output.OutDate{Time: todo.DueDate()},
		Priority:           todo.PriorityColor(),
		CompleteFlag:       todo.CompleteFlag(),
	}, nil
}
