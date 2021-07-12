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
