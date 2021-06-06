package datasource

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Todo struct {
	ID                 int       `db:"id"`
	Title              string    `db:"title"`
	Memo               *string   `db:"memo"`
	ImplementationDate time.Time `db:"implementation_date"`
	DueDate            time.Time `db:"due_date"`
	PriorityColor      string    `db:"color"`
	CompleteFlag       BitBool   `db:"complete_flag"`
}

type BitBool bool

func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	}
	return []byte{0}, nil
}

func (b *BitBool) Scan(src interface{}) error {
	bitBool, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = bitBool[0] == 1
	return nil
}
