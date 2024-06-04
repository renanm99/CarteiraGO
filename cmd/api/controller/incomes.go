package controller

import (
	"carteirago/cmd/api/db/repository"
	"carteirago/cmd/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IncomesGET(c *gin.Context) {
	route := "incomes"
	method := c.Request.Method
	id := c.DefaultQuery("id", "")
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, incomes, err := repository.IncomesSelect(userid)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if len(incomes) == 0 {
		c.JSON(204, gin.H{"route": route, "method": method, "userid": userid, "incomes": ""})
		return
	}

	if id == "" {
		c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "incomes": incomes})
		return
	}
}

func IncomesPOST(c *gin.Context) {
	route := "incomes"
	method := c.Request.Method
	userid, err := strconv.Atoi(c.Query("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "incomes": income})
}

func IncomesPUT(c *gin.Context) {
	route := "incomes"
	method := c.Request.Method
	userid, err := strconv.Atoi(c.Query("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "incomes": income})
}

func IncomesDELETE(c *gin.Context) {
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

	code, err := repository.IncomesDelete(userid, id)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
	c.Done()
}
