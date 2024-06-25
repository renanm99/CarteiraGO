package repository

import (
	"carteirago/api/db"
	"carteirago/api/models"
	"crypto/sha256"
	"fmt"
	"net/http"
)

func CustomerSelect(userid int) (int, *models.Customer, error) {
	dbConn, _ := db.Database()

	query := fmt.Sprintf("select * from customer where user_id = %d", userid)
	row := dbConn.QueryRow(query)

	customer := new(models.Customer)
	if err := row.Scan(&customer.Id, &customer.Fullname, &customer.Email, &customer.Password, &customer.Socialname); err != nil {
		return http.StatusInternalServerError, customer, err
	}

	dbConn.Close()

	return http.StatusOK, customer, nil
}

func CustomerCheck(email string) int {
	dbConn, _ := db.Database()

	query := fmt.Sprintf("select id from customer where email = '%s'", email)
	row := dbConn.QueryRow(query)

	var Id int
	if err := row.Scan(&Id); err != nil {
		Id = 0
	}

	dbConn.Close()

	return Id
}

func CustomerInsert(customer *models.Customer) (int, error) {

	dbConn, _ := db.Database()
	query := fmt.Sprintf("insert into customer (fullname,	email,	password,	socialname)"+
		" values ('%s','%s','%s','%s')",
		customer.Fullname,
		customer.Email,
		HashPwd(customer.Password+customer.Email),
		customer.Socialname)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusOK, nil
}

func CustomerUpdate(customer *models.Customer) (int, error) {
	dbConn, _ := db.Database()
	query := fmt.Sprintf("update customer set fullname = '%s', email = '%s', password = '%s', socialname = %s, "+
		"where id = %d",
		customer.Fullname,
		customer.Email,
		HashPwd(customer.Password+customer.Email),
		customer.Socialname,
		customer.Id)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusOK, nil
}

func CustomerDelete(userid int) (int, error) {
	dbConn, _ := db.Database()
	query := fmt.Sprintf("delete from expenses where user_id = %d; "+
		"delete from incomes where user_id = %d; "+
		"delete from customer where id = %d;",
		userid, userid, userid)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusNoContent, nil
}

func HashPwd(pwd string) string {
	h := sha256.New()
	h.Write([]byte(pwd))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func CheckPwd(email string, pwd string) int {
	dbConn, _ := db.Database()

	h := sha256.New()
	pass := fmt.Sprintf("%s%s", pwd, email)
	h.Write([]byte(pass))
	check := fmt.Sprintf("%x", h.Sum(nil))
	query := fmt.Sprintf("select id from customer where email = '%s' and password = '%s'", email, check)

	var id int
	if err := dbConn.QueryRow(query).Scan(&id); err != nil {
		id = 0
	}

	dbConn.Close()
	return id
}
