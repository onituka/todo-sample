package usecase

import (
	"github.com/onituka/todo-sample/domain/tododomain"
	"github.com/onituka/todo-sample/usecase/input"
	"github.com/onituka/todo-sample/usecase/output"
)

type TodoUsecase interface {
	FetchTodo(todoID int) (*output.Todo, error)
	FetchAllTodo() ([]output.Todo, error)
	Create(in *input.Todo) (*output.Todo, error)
	Update(up *input.Todo) (*output.Todo, error)
	Delete(todoID int) error
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

//全件取得
func (u todoUsecase) FetchAllTodo() ([]output.Todo, error) {
	todos, err := u.todoRepository.FetchAllTodo()
	if err != nil {
		return nil, err
	}

	todosDto := make([]output.Todo, len(todos))

	for i, todo := range todos {
		todosDto[i] = output.Todo{
			ID:                 todo.ID(),
			Title:              todo.Title(),
			Memo:               todo.Memo(),
			ImplementationDate: output.OutDate{Time: todo.ImplementationDate()},
			DueDate:            output.OutDate{Time: todo.DueDate()},
			Priority:           todo.PriorityColor(),
			CompleteFlag:       todo.CompleteFlag(),
		}

	}

	return todosDto, nil
}

//新規作成
func (u *todoUsecase) Create(in *input.Todo) (*output.Todo, error) {
	todo := tododomain.NewAddTodo(in.Title, in.Memo, in.ImplementationDate.Time, in.DueDate.Time, in.Priority, in.CompleteFlag)
	id, err := u.todoRepository.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	//id を使用してselectする
	todo, err = u.todoRepository.FetchTodo(id)
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

//更新
func (u *todoUsecase) Update(up *input.Todo) (*output.Todo, error) {
	todo := tododomain.NewTodo(up.ID, up.Title, up.Memo, up.ImplementationDate.Time, up.DueDate.Time, up.Priority, up.CompleteFlag)
	if err := u.todoRepository.UpdateTodo(todo); err != nil {
		return nil, err
	}

	todo, err := u.todoRepository.FetchTodo(todo.ID())
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

//削除
func (u *todoUsecase) Delete(todoID int) error {
	if err := u.todoRepository.DeleteTodo(todoID); err != nil {
		return err
	}

	return nil
}
