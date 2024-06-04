package models

type Customer struct {
	Id         int    `db: "id" json: "id"`
	Fullname   string `db: "fullname" json: "fullname"`
	Email      string `db: "email" json: "email"`
	Password   string `db: "password" json: "password"`
	Socialname string `db: "socialname" json: "socialname"`
}
