package repository

import (
	"carteirago/cmd/api/db"
	"carteirago/cmd/api/models"
	"fmt"
	"net/http"
	"time"
)

func IncomesSelect(userid int) (int, []models.Incomes, error) {
	dbConn := db.Database()

	incomes := []models.Incomes{}
	query := fmt.Sprintf("select * from incomes where user_id = %d", userid)
	rows, err := dbConn.Query(query)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	defer rows.Close()

	for rows.Next() {
		income := new(models.Incomes)
		if err := rows.Scan(&income.Id, &income.UserId, &income.Title, &income.Description, &income.Type, &income.Value, &income.Datetime); err != nil {
			return http.StatusInternalServerError, nil, err
		}
		incomes = append(incomes, *income)
	}

	dbConn.Close()

	return http.StatusOK, incomes, nil
}

func IncomesInsert(income *models.Incomes) (int, error) {
	dbConn := db.Database()
	query := fmt.Sprintf("insert into incomes (user_id,	title,	description,	type,	value,	datetime)"+
		" values (%d,'%s','%s','%s',%f,'%s')",
		income.UserId,
		income.Title, income.Description,
		income.Type,
		income.Value,
		income.Datetime.Format(time.DateTime))

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusOK, nil
}

func IncomesUpdate(income *models.Incomes) (int, error) {
	dbConn := db.Database()
	query := fmt.Sprintf("update incomes set title = '%s', description = '%s', type = '%s', value = %f, "+
		"datetime = '%s' where userid = %d and id = %d",
		income.Title, income.Description,
		income.Type,
		income.Value,
		income.Datetime.Format(time.DateTime),
		income.UserId, income.Id)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusNoContent, nil
}

func IncomesDelete(userid int, incomeid int) (int, error) {
	dbConn := db.Database()
	query := fmt.Sprintf("delete from incomes where userid = %d and id = %d", userid, incomeid)

	_, err := dbConn.Exec(query)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	dbConn.Close()

	return http.StatusNoContent, nil
}
