package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/onituka/todo-sample/apierrors"
	"github.com/onituka/todo-sample/interfaces/presenter"
	"github.com/onituka/todo-sample/usecase"
	"github.com/onituka/todo-sample/usecase/input"
)

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) *todoHandler {
	return &todoHandler{
		todoUsecase: todoUsecase,
	}
}

func (h *todoHandler) FetchTodo(w http.ResponseWriter, r *http.Request) {
	todoID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		presenter.ErrorJSON(w, apierrors.NewBadRequestError(apierrors.NewErrorString("todo idを正しく入力してください")))
	}

	out, err := h.todoUsecase.FetchTodo(todoID)
	if err != nil {
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}

//全件取得
func (h *todoHandler) FetchAllTodo(w http.ResponseWriter, r *http.Request) {
	items, err := h.todoUsecase.FetchAllTodo()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err = json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), 500)
	}

}

//新規作成
func (h todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var in input.Todo
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		return
	}

	out, err := h.todoUsecase.Create(&in)
	if err != nil {
		presenter.ErrorJSON(w, apierrors.NewInternalServerError(apierrors.NewErrorString("Invalid usecase")))
		return
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
