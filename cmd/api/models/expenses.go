package models

import (
	"time"
)

type Expenses struct {
	Id          int       `db: "id"	json: "Id"`
	UserId      int       `db: "user_id"	json: "UserId" binding:"required"`
	Title       string    `db: "title"	json: "Title" binding:"required"`
	Description string    `db: "description"	json: "Description" binding:"required"`
	Type        string    `db: "type"	json: "Type" binding:"required"`
	Value       float32   `db: "value"	json: "Value binding:"required"`
	Datetime    time.Time `db: "datetime"	json: "Datetime"`
}
