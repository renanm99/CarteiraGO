package controller

import (
	"carteirago/cmd/api/db/repository"
	"carteirago/cmd/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IncomesGET(c *gin.Context) {

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

	code, incomes, err := repository.IncomesSelect(userid)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if len(incomes) == 0 {
		c.JSON(http.StatusOK, gin.H{"incomes": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"incomes": incomes})
	c.Done()
}

func IncomesPOST(c *gin.Context) {

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

	income := new(models.Incomes)

	err = c.ShouldBindJSON(&income)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	income.UserId = userid

	code, err := repository.IncomesInsert(income)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.String(200, "")
	c.Done()
	return
}

func IncomesPUT(c *gin.Context) {
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

	income := new(models.Incomes)

	err = c.ShouldBindJSON(&income)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	income.UserId = userid

	code, err := repository.IncomesUpdate(income)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "")
	c.Done()
	return
}

func IncomesDELETE(c *gin.Context) {
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

	var id int
	err = c.ShouldBindJSON(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code, err := repository.IncomesDelete(userid, id)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
	c.Done()
	return
}
