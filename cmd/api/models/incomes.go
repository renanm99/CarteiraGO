package models

import "time"

type Incomes struct {
	Id          int64     `json: "id" binding:"required"`
	UserId      int32     `json: "userid" binding:"required"`
	Title       string    `json: "title" binding:"required"`
	Description string    `json: "description" binding:"required"`
	Type        string    `json: "type" binding:"required"`
	Value       float32   `json: "value" binding:"required"`
	Datetime    time.Time `json: "datetime" binding:"required"`
	CreatedAt   time.Time `json: "createdAt"`
}
