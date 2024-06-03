package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "aws-0-sa-east-1.pooler.supabase.com"
	port   = "6543"
	user   = "postgres.upglzprsczdxnepvajzm"
	dbname = "db_carteirago"
)

func Database() *sql.DB {
	host, user, password, env, dbname := DatabaseConstants()
	psqlInfo := fmt.Sprintf("host=%s\nuser=%s\n"+
		"password=%s\ndbname=%s\nsslmode=disable",
		host, user, password, dbname)

	if !env {
		panic("os env password empty")
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//fmt.Println(psqlInfo)

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("\ndb connected!")

	return db
}

func DatabaseConstants() (string, string, string, bool, string) {
	key, env := os.LookupEnv("SUPABASE_KEY")
	return host, user, key, env, dbname
}
