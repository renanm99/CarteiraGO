package models

import "time"

type User struct {
	Id        int32     `json: "id" binding:"required"`
	Name      string    `json: "name" binding:"required"`
	Age       int       `json:"age" binding:"required"`
	CreatedAt time.Time `json: "createdAt"`
}
