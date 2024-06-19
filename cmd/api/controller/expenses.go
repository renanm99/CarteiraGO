package controller

import (
	"carteirago/cmd/api/db/repository"
	"carteirago/cmd/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExpensesGET(c *gin.Context) {
	id, token, err := getCookieHandler(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if err := verifyToken(token); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	userid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code, expenses, err := repository.ExpensesSelect(userid)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if len(expenses) == 0 {
		c.JSON(http.StatusOK, gin.H{"expenses": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
	c.Done()
}

func ExpensesPOST(c *gin.Context) {
	id, token, err := getCookieHandler(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if err := verifyToken(token); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	userid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	c.String(http.StatusOK, "")
	c.Done()
	return
}

func ExpensesPUT(c *gin.Context) {
	id, token, err := getCookieHandler(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if err := verifyToken(token); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	userid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expense := new(models.Expenses)

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

	c.String(http.StatusOK, "")
	c.Done()
	return
}

func ExpensesDELETE(c *gin.Context) {

	user, token, err := getCookieHandler(c)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if err := verifyToken(token); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	userid, err := strconv.Atoi(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expense := new(models.Expenses)
	err = c.ShouldBindJSON(&expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code, err := repository.ExpensesDelete(userid, expense.Id)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
	c.Done()
	return
}
