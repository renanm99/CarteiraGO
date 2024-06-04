package repository

import (
	"carteirago/cmd/api/db"
	"carteirago/cmd/api/models"
	"fmt"
	"net/http"
	"time"
)

func ExpensesSelect(userid int) (int, []models.Expenses, error) {
	dbConn := db.Database()

	expenses := []models.Expenses{}
	query := fmt.Sprintf("select * from expenses where user_id = %d", userid)
	rows, err := dbConn.Query(query)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		expense := new(models.Expenses)
		if err := rows.Scan(&expense.Id, &expense.UserId, &expense.Title, &expense.Description, &expense.Type, &expense.Value, &expense.Datetime); err != nil {
			return http.StatusInternalServerError, nil, err
		}
		expenses = append(expenses, *expense)
	}

	dbConn.Close()

	return http.StatusOK, expenses, nil
}

func ExpensesInsert(expense *models.Expenses) (int, error) {
	dbConn := db.Database()
	query := fmt.Sprintf("insert into expenses (user_id,	title,	description,	type,	value,	datetime)"+
		" values (%d,'%s','%s','%s',%f,'%s')",
		expense.UserId,
		expense.Title, expense.Description,
		expense.Type,
		expense.Value,
		expense.Datetime.Format(time.DateTime))

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusOK, nil
}

func ExpensesUpdate(expense *models.Expenses) (int, error) {
	dbConn := db.Database()
	query := fmt.Sprintf("update expenses set title = '%s', description = '%s', type = '%s', value = %f, "+
		"datetime = '%s' where userid = %d and id = %d",
		expense.Title, expense.Description,
		expense.Type,
		expense.Value,
		expense.Datetime.Format(time.DateTime),
		expense.UserId, expense.Id)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusNoContent, nil
}

func ExpensesDelete(userid int, expenseid int) (int, error) {
	dbConn := db.Database()
	query := fmt.Sprintf("delete from expenses where userid = %d and id = %d", userid, expenseid)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusNoContent, nil
}
