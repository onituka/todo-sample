package input

import "time"

type Todo struct {
	ID                 int       `json:"id"`
	Title              string    `json:"title"`
	Memo               string    `json:"memo"`
	ImplementationDate time.Time `json:"implementation_date"`
	DueDate            time.Time `json:"due_date"`
	Priority           string    `json:"priority"`
	CompleteFlag       bool      `json:"complete_flag"`
}
