package input

import (
	"strings"
	"time"
)

type Todo struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Memo               *string `json:"memo"`
	ImplementationDate InDate  `json:"implementation_date"`
	DueDate            InDate  `json:"due_date"`
	Priority           string  `json:"priority"`
	CompleteFlag       bool    `json:"complete_flag"`
}

type InDate struct {
	time.Time
}

func (d *InDate) UnmarshalJSON(data []byte) error {
	strDate := strings.Trim(string(data), "\"")
	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		return err
	}

	d.Time = date

	return nil
}
