package models

import "time"

type Expenses struct {
	Id          int64     `json: "id" `
	UserId      int32     `json: "userid" `
	Title       string    `json: "title" `
	Description string    `json: "description" `
	Type        string    `json: "type" `
	Value       float32   `json: "value" `
	Datetime    time.Time `json: "datetime" `
	CreatedAt   time.Time `json: "createdAt"`
}
