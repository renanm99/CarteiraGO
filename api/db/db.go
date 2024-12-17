package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   string = "aws-0-sa-east-1.pooler.supabase.com"
	port   string = "6543"
	user   string = "postgres.upglzprsczdxnepvajzm"
	dbname string = "db_carteirago"
)

func Database() (*sql.DB, error) {
	host, user, password, env, dbname := DatabaseConstants()
	psqlInfo := fmt.Sprintf("host=%s\nuser=%s\n"+
		"password=%s\ndbname=%s\nsslmode=disable",
		host, user, password, dbname)

	if !env {
		//panic("os env password empty")
		return nil, errors.New("os env password empty")
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	//fmt.Println(psqlInfo)

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func DatabaseConstants() (string, string, string, bool, string) {
	key, env := os.LookupEnv("supabase_key")
	return host, user, key, env, dbname
}
