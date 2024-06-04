package models

type Customer struct {
	Id         int32  `db: "id" json: "id"`
	Fullname   string `db: "user_id" json: "user_id"`
	Email      string `db: "title" json: "title"`
	Password   string `db: "description" json: "description"`
	Socialname string `db: "type" json: "type"`
}
