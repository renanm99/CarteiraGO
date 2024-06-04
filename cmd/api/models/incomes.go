package models

import "time"

type Incomes struct {
	Id          int64     `db: "id"	json: "id"`
	UserId      int32     `db: "user_id"	json: "user_id"`
	Title       string    `db: "title"	json: "title"`
	Description string    `db: "description"	json: "description"`
	Type        string    `db: "type"	json: "type"`
	Value       float32   `db: "value"	json: "value"`
	Datetime    time.Time `db: "datetime"	json: "datetime"`
}
