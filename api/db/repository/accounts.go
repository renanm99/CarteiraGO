package repository

import (
	"carteirago/api/db"
	"carteirago/api/models"
	"fmt"
	"net/http"
)

func AccountsSelect(userid int, account string) (int, []models.Accounts, error) {
	dbConn, _ := db.Database()

	accounts := []models.Accounts{}
	query := fmt.Sprintf("select * from %s where user_id = %d order by id", account, userid)
	rows, err := dbConn.Query(query)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	defer rows.Close()
	defer dbConn.Close()

	for rows.Next() {
		account := new(models.Accounts)
		if err := rows.Scan(&account.Id, &account.UserId, &account.Title, &account.Description, &account.Category, &account.Value, &account.Datetime); err != nil {
			return http.StatusInternalServerError, nil, err
		}
		accounts = append(accounts, *account)
	}

	dbConn.Close()

	return http.StatusOK, accounts, nil
}

func AccountsInsert(account *models.Accounts) (int, error) {
	dbConn, _ := db.Database()
	query := fmt.Sprintf("insert into %s (user_id,title,description,category,value,datetime)"+
		" values (%d,'%s','%s','%s',%f, NOW()::timestamp)",
		account.Account,
		account.UserId,
		account.Title, account.Description,
		account.Category,
		account.Value)

	_, err := dbConn.Exec(query)

	defer dbConn.Close()

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusOK, nil
}

func AccountsUpdate(account *models.Accounts) (int, error) {
	dbConn, _ := db.Database()
	query := fmt.Sprintf("update %s set title='%s',description='%s',category='%s',value=%f"+
		" where user_id=%d and id = %d",
		account.Account,
		account.Title, account.Description,
		account.Category,
		account.Value,
		account.UserId, account.Id)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusNoContent, nil
}

func AccountsDelete(userid int, accountid int, account string) (int, error) {
	dbConn, _ := db.Database()
	query := fmt.Sprintf("delete from %s where user_id = %d and id = %d", account, userid, accountid)

	_, err := dbConn.Exec(query)

	defer dbConn.Close()

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusNoContent, nil
}

func DashboardSelect(userid int, account string) (int, []models.Dashboard, error) {
	dbConn, _ := db.Database()

	accounts := []models.Dashboard{}
	query := fmt.Sprintf("select category, value, datetime from %s WHERE user_id = %d", account, userid)
	rows, err := dbConn.Query(query)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	defer rows.Close()
	defer dbConn.Close()

	for rows.Next() {
		account := new(models.Dashboard)
		if err := rows.Scan(&account.Category, &account.Value, &account.Datetime); err != nil {
			return http.StatusInternalServerError, nil, err
		}
		accounts = append(accounts, *account)
	}

	dbConn.Close()

	return http.StatusOK, accounts, nil
}
