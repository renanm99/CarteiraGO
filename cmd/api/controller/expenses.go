package controller

import (
	"carteirago/cmd/api/db/repository"
	"carteirago/cmd/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExpensesGET(c *gin.Context) {
	route := "expenses"
	method := c.Request.Method
	/*
		userid, err := strconv.Atoi(c.Query("userid"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	*/
	user, token, err := getCookieHandler(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if err := verifyToken(token); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	code, expenses, err := repository.ExpensesSelect(user)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if len(expenses) == 0 {
		c.JSON(http.StatusOK, gin.H{"route": route, "method": method, "expenses": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"route": route, "method": method, "expenses": expenses})
	c.Done()
	return
}

func ExpensesPOST(c *gin.Context) {
	route := "expenses"
	method := c.Request.Method
	userid, err := strconv.Atoi(c.Query("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense := new(models.Expenses)
	err = c.ShouldBindJSON(&expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expense.UserId = userid

	code, err := repository.ExpensesInsert(expense)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"route": route, "method": method, "userid": userid, "expenses": expense})
	c.Done()
	return
}

func ExpensesPUT(c *gin.Context) {
	route := "expenses"
	method := c.Request.Method
	userid, err := strconv.Atoi(c.Query("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense := new(models.Expenses)
	expense.UserId = userid
	err = c.ShouldBindJSON(&expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expense.UserId = userid

	code, err := repository.ExpensesUpdate(expense)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"route": route, "method": method, "userid": userid, "expenses": expense})
	c.Done()
	return
}

func ExpensesDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userid, err := strconv.Atoi(c.Query("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, err := repository.ExpensesDelete(userid, id)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
	c.Done()
	return
}
