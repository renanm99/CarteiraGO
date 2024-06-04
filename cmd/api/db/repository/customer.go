package repository

import (
	"carteirago/cmd/api/db"
	"carteirago/cmd/api/models"
	"fmt"
	"net/http"
)

func CustomerGET(userid int) (int, *models.Customer, error) {
	dbConn := db.Database()

	query := fmt.Sprintf("select * from customer where user_id = %d", userid)
	row := dbConn.QueryRow(query)

	customer := new(models.Customer)
	if err := row.Scan(&customer.Id, &customer.Fullname, &customer.Email, &customer.Password, &customer.Socialname); err != nil {
		return http.StatusInternalServerError, customer, err
	}

	dbConn.Close()

	return http.StatusOK, customer, nil
}

func CustomerPOST(customer *models.Customer) (int, error) {
	dbConn := db.Database()
	query := fmt.Sprintf("insert into customer (fullname,	email,	password,	socialname)"+
		" values ('%s','%s','%s','%s')",
		customer.Fullname,
		customer.Email,
		customer.Password,
		customer.Socialname)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusOK, nil
}
