package tododomain

import "time"

type Todo struct {
	id                 int
	title              string
	memo               *string
	implementationDate time.Time
	dueDate            time.Time
	priorityColor      string
	completeFlag       bool
}

func NewTodo(id int, title string, memo *string, implementationDate time.Time, dueDate time.Time, priorityColor string, completeFlag bool) *Todo {
	return &Todo{
		id:                 id,
		title:              title,
		memo:               memo,
		implementationDate: implementationDate,
		dueDate:            dueDate,
		priorityColor:      priorityColor,
		completeFlag:       completeFlag,
	}
}

func NewAddTodo(title string, memo *string, implementationDate time.Time, dueDate time.Time, priorityColor string, completeFlag bool) *Todo {
	return &Todo{
		title:              title,
		memo:               memo,
		implementationDate: implementationDate,
		dueDate:            dueDate,
		priorityColor:      priorityColor,
		completeFlag:       completeFlag,
	}

}

func (t *Todo) ID() int {
	return t.id
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) Memo() *string {
	return t.memo
}

func (t *Todo) ImplementationDate() time.Time {
	return t.implementationDate
}

func (t *Todo) DueDate() time.Time {
	return t.dueDate
}

func (t *Todo) PriorityColor() string {
	return t.priorityColor
}

func (t *Todo) CompleteFlag() bool {
	return t.completeFlag
}
