package controller

import (
	"carteirago/cmd/api/db/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExpensesGET(c *gin.Context) {
	route := "ExpensesGET"
	method := c.Request.Method
	id := c.DefaultQuery("id", "")
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	expenses := repository.ExpensesGET(userid)

	if len(expenses) == 0 {
		c.JSON(204, gin.H{"route": route, "method": method, "userid": userid, "incomes": ""})
		return
	}

	if id == "" {
		c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "expenses": expenses})
		return
	}
}

func ExpensesPOST(c *gin.Context) {
	route := "expenses"
	method := c.Request.Method
	userid := c.Query("userid")
	id := c.Query("id")
	c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "expenses": id})
}
