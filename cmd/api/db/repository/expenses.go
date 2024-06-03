package repository

import (
	"carteirago/cmd/api/db"
	"carteirago/cmd/api/models"
	"fmt"
)

func ExpensesGET(userid int) []models.Expenses {
	dbConn := db.Database()

	expenses := []models.Expenses{}
	query := fmt.Sprintf("select * from expenses where user_id = %d", userid)
	rows, err := dbConn.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		expense := new(models.Expenses)
		if err := rows.Scan(&expense.Id, &expense.UserId, &expense.Title, &expense.Description, &expense.Type, &expense.Value, &expense.Datetime); err != nil {
			panic(err)
		}
		expenses = append(expenses, *expense)
	}

	dbConn.Close()

	return expenses
}

func ExpensesPOST() {
}
