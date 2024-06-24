package models

import (
	"time"
)

type Accounts struct {
	Id          int       `db: "id"	json: "Id"`
	UserId      int       `db: "user_id"	json: "UserId" binding:"required"`
	Title       string    `db: "title"	json: "Title" binding:"required"`
	Description string    `db: "description"	json: "Description" binding:"required"`
	Category    string    `db: "category"	json: "Category" binding:"required"`
	Value       float32   `db: "value"	json: "Value binding:"required"`
	Datetime    time.Time `db: "datetime"	json: "Datetime"`
	Account     string    `json: "Account"`
}

type Dashboard struct {
	Category string    `db: "category"`
	Value    float32   `db: "value"`
	Datetime time.Time `db: datetime`
}
