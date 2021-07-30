package usecase

import (
	"testing"
	"time"

	mock_tododomain "github.com/onituka/todo-sample/usecase/mock_repository"

	"github.com/onituka/todo-sample/usecase/input"

	"github.com/golang/mock/gomock"
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
func (m *mockTodoRepository) FetchAllTodo() ([]*tododomain.Todo, error) {
	memo := "testです"
	memo2 := "testです"

	return []*tododomain.Todo{
		tododomain.NewTodo(
			1,
			"test todo",
			&memo,
			time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
			time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
			"#ff0000",
			true,
		),
		tododomain.NewTodo(
			2,
			"test todo",
			&memo2,
			time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
			time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
			"#ff0000",
			true,
		),
	}, nil
}

func (m *mockTodoRepository) CreateTodo(todo *tododomain.Todo) (int, error) {
	return 1, nil
}

func (m *mockTodoRepository) UpdateTodo(todo *tododomain.Todo) error {
	return nil
}

func (m *mockTodoRepository) DeleteTodo(todoID int) error {
	return nil
}

func Test_todoUsecase_FetchTodo(t *testing.T) {
	u := NewTodoUsecase(&mockTodoRepository{})

	in := 1

	gotOut, err := u.FetchTodo(in)
	if err != nil {
		t.Errorf("unexpected error by todoUsecase.FetchTodo '%#v'", err)
		return
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

func TestTodoUsecase_FetchAllTodo(t *testing.T) {
	u := NewTodoUsecase(&mockTodoRepository{})

	gotTodos, err := u.FetchAllTodo()
	if err != nil {
		t.Errorf("unexpected error by todoUsecase.FetchTodo '%#v'", err)
		return
	}

	wantMemo := "testです"
	wantMemo2 := "testです"

	wantOut := []output.Todo{
		{
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
		},
		{
			ID:    2,
			Title: "test todo",
			Memo:  &wantMemo2,
			ImplementationDate: output.OutDate{
				Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
			},
			DueDate: output.OutDate{
				Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
			},
			Priority:     "#ff0000",
			CompleteFlag: true,
		},
	}

	if diff := cmp.Diff(&wantOut, &gotTodos); len(diff) != 0 {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

func TestTodoUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreate := mock_tododomain.NewMockRepository(ctrl)
	wantmemo := "testです"
	todo := tododomain.NewAddTodo(
		"test",
		&wantmemo,
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		"#ff0000",
		true,
	)
	todo2 := tododomain.NewTodo(
		1,
		"test",
		&wantmemo,
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		"#ff0000",
		true,
	)

	mockCreate.EXPECT().CreateTodo(todo).Return(1, nil)
	mockCreate.EXPECT().FetchTodo(1).Return(todo2, nil)

	u := NewTodoUsecase(mockCreate)
	in := &input.Todo{
		ID:    1,
		Title: "test",
		Memo:  &wantmemo,
		ImplementationDate: input.InDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		DueDate: input.InDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		Priority:     "#ff0000",
		CompleteFlag: true,
	}
	gotOut, err := u.Create(in)
	if err != nil {
		t.Errorf("unexpected error by todoUsecase.FetchTodo '%#v'", err)
		return
	}

	wantOut := &output.Todo{
		ID:    1,
		Title: "test",
		Memo:  &wantmemo,
		ImplementationDate: output.OutDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		DueDate: output.OutDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		Priority:     "#ff0000",
		CompleteFlag: true,
	}

	if diff := cmp.Diff(gotOut, wantOut); len(diff) != 0 {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}

}

func TestTodoUsecase_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUpdate := mock_tododomain.NewMockRepository(ctrl)
	wantmemo := "testです"
	todo := tododomain.NewTodo(
		1,
		"test",
		&wantmemo,
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		"#ff0000",
		true,
	)

	todo2 := tododomain.NewTodo(
		1,
		"test",
		&wantmemo,
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		"#ff0000",
		true,
	)
	mockUpdate.EXPECT().UpdateTodo(todo).Return(nil)
	mockUpdate.EXPECT().FetchTodo(1).Return(todo2, nil)

	u := NewTodoUsecase(mockUpdate)

	in := &input.Todo{
		ID:    1,
		Title: "test",
		Memo:  &wantmemo,
		ImplementationDate: input.InDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		DueDate: input.InDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		Priority:     "#ff0000",
		CompleteFlag: true,
	}
	gotOut, err := u.Update(in)
	if err != nil {
		t.Errorf("unexpected error by todoUsecase.FetchTodo '%#v'", err)
		return
	}

	wantOut := &output.Todo{
		ID:    1,
		Title: "test",
		Memo:  &wantmemo,
		ImplementationDate: output.OutDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		DueDate: output.OutDate{
			Time: time.Date(2021, 6, 20, 0, 0, 0, 0, time.Local),
		},
		Priority:     "#ff0000",
		CompleteFlag: true,
	}

	if diff := cmp.Diff(gotOut, wantOut); len(diff) != 0 {
		t.Errorf("differs: (-want +got)\n%s", diff)
	}
}

func TestTodoUsecase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDelete := mock_tododomain.NewMockRepository(ctrl)

	mockDelete.EXPECT().DeleteTodo(1).Return(nil)
	t.Log("result:", mockDelete.DeleteTodo(1))

}
