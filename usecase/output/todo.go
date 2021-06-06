package output

import (
	"fmt"
	"time"
)

type Todo struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Memo               *string `json:"memo"`
	ImplementationDate OutDate `json:"implementation_date"`
	DueDate            OutDate `json:"due_date"`
	Priority           string  `json:"priority"`
	CompleteFlag       bool    `json:"complete_flag"`
}

type OutDate struct {
	time.Time
}

func (d *OutDate) MarshalJSON() ([]byte, error) {
	date := d.Time.Format("2006-01-02")

	return []byte(fmt.Sprintf("\"%s\"", date)), nil
}
