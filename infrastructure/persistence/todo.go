package persistence

import (
	"database/sql"

	"golang.org/x/xerrors"

	"github.com/onituka/todo-sample/apierrors"
	"github.com/onituka/todo-sample/domain/tododomain"
	"github.com/onituka/todo-sample/infrastructure/persistence/datasource"
	"github.com/onituka/todo-sample/infrastructure/persistence/rdb"
)

type todoRepository struct {
	*rdb.MySQLHandler
}

func NewTodoRepository(mysqlHandler *rdb.MySQLHandler) *todoRepository {
	return &todoRepository{mysqlHandler}
}

func (r *todoRepository) FetchTodo(todoID int) (*tododomain.Todo, error) {
	query := `
        SELECT
            todos.id                  id,
            todos.title               title,
            todos.memo                memo,
            todos.implementation_date implementation_date,
            todos.due_date            due_date,
            priorities.color          color,
            todos.complete_flag       complete_flag
        FROM
            todos
        INNER JOIN
            priorities
        ON
            priorities.id = todos.priorities_id
        WHERE
            todos.id = ?`

	var todoDto datasource.Todo
	if err := r.MySQLHandler.Conn.QueryRowx(query, todoID).StructScan(&todoDto); err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return nil, apierrors.NewNotFoundError(apierrors.NewErrorString("指定されたtodoは存在しません"))
		}

		return nil, apierrors.NewInternalServerError(apierrors.NewErrorString("Internal Server Error"))
	}

	todo := tododomain.NewTodo(todoDto.ID, todoDto.Title, todoDto.Memo, todoDto.ImplementationDate, todoDto.DueDate, todoDto.PriorityColor, bool(todoDto.CompleteFlag))

	return todo, nil
}

//全件取得
func (r *todoRepository) FetchAllTodo() ([]*tododomain.Todo, error) {
	query := `
        SELECT
            todos.id                  id,
            todos.title               title,
            todos.memo                memo,
            todos.implementation_date implementation_date,
            todos.due_date            due_date,
            priorities.color          color,
            todos.complete_flag       complete_flag
        FROM
            todos
        INNER JOIN
            priorities
        ON
            priorities.id = todos.priorities_id`
	rows, err := r.MySQLHandler.Conn.Queryx(query)
	if err != nil {
		return nil, apierrors.NewInternalServerError(apierrors.NewErrorString("Internal Server error"))
	}

	defer rows.Close()
	var todosDto []datasource.Todo
	for rows.Next() {
		var todoDto datasource.Todo
		if err := rows.StructScan(&todoDto); err != nil {
			return nil, apierrors.NewInternalServerError(apierrors.NewErrorString("Internal Server Error"))
		}

		todosDto = append(todosDto, todoDto)
	}

	todos := make([]*tododomain.Todo, len(todosDto))
	for i, todoDto := range todosDto {
		todo := tododomain.NewTodo(todoDto.ID, todoDto.Title, todoDto.Memo, todoDto.ImplementationDate, todoDto.DueDate, todoDto.PriorityColor, bool(todoDto.CompleteFlag))
		todos[i] = todo
	}
	return todos, nil
}

//新規作成
func (r *todoRepository) CreateTodo(todo *tododomain.Todo) (int, error) {
	query := `
         INSERT INTO todos
        (
            title,
            memo,
            implementation_date,
            due_date,
            priorities_id,
            complete_flag
        )
         VALUES
            (?, ?, ?, ?, ?, ?)`

	result, err := r.MySQLHandler.Conn.Exec(query, todo.Title(), todo.Memo(), todo.ImplementationDate(), todo.DueDate(), 1, todo.CompleteFlag())
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return int(id), nil
}

//更新
func (r *todoRepository) UpdateTodo(todo *tododomain.Todo) error {
	query := `
          UPDATE 
               todos
          SET
               title = ?,
               memo = ?,
               implementation_date = ?,
               due_date = ?,
               priorities_id = ?,
               complete_flag = ?
           WHERE
            todos.id = ?`

	_, err := r.MySQLHandler.Conn.Exec(query, todo.Title(), todo.Memo(), todo.ImplementationDate(), todo.DueDate(), 1, todo.CompleteFlag(), todo.ID())
	if err != nil {
		return apierrors.NewInternalServerError(apierrors.NewErrorString("Internal Server Error"))
	}

	return err
}

//削除
func (r *todoRepository) DeleteTodo(todoID int) error {
	query := `
           DELETE
           FROM
               todos
           WHERE
	           todos.id = ?`

	result, err := r.MySQLHandler.Conn.Exec(query, todoID)
	if err != nil {
		return apierrors.NewInternalServerError(apierrors.NewErrorString("Internal Server Error"))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apierrors.NewInternalServerError(apierrors.NewErrorString("Internal Server Error"))
	}

	if rowsAffected == 0 {
		return apierrors.NewNotFoundError(apierrors.NewErrorString("todoはすでに削除されています"))
	}
	return nil
}
