package models

type Customer struct {
	Id         int    `db: "id" json: "Id"`
	Fullname   string `db: "fullname" json: "Fullname"`
	Email      string `db: "email" json: "Email"`
	Password   string `db: "password" json: "Password"`
	Socialname string `db: "socialname" json: "Socialname"`
}
