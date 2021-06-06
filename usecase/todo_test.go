package usecase

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/onituka/todo-sample/domain/tododomain"
	"github.com/onituka/todo-sample/usecase/output"
)

type mockTodoRepository struct{}

func (m *mockTodoRepository) FetchTodo(todoID int) (*tododomain.Todo, error) {
	memo := "testです"
	return tododomain.NewTodo(
		1,
		"test todo",
		&memo,
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		"#ff0000",
		true,
	), nil
}

func Test_todoUsecase_FetchTodo(t *testing.T) {
	u := NewTodoUsecase(&mockTodoRepository{})

	in := 1

	gotOut, err := u.FetchTodo(in)
	if err != nil {
		t.Errorf("unexpected error by todoUsecase.FetchTodo '%#v'", err)
	}

	wantMemo := "testです"

	wantOut := &output.Todo{
		ID:    1,
		Title: "test todo",
		Memo:  &wantMemo,
		ImplementationDate: output.OutDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		DueDate: output.OutDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		Priority:     "#ff0000",
		CompleteFlag: true,
	}

	if diff := cmp.Diff(&wantOut, &gotOut); len(diff) != 0 {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}
