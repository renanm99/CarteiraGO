package controller

import (
	"carteirago/cmd/api/models"
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

//const repo string = "C:\\dev\\carteirago\\cmd\\api\\db\\repo_data"

func ExpensesGET(c *gin.Context) {
	route := "ExpensesGET"
	method := c.Request.Method
	id := c.DefaultQuery("id", "")
	userid, err := strconv.ParseInt(c.Query("userid"), 10, 8)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonFile, err := os.Open(repo + "/expenses.json")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonParser := json.NewDecoder(jsonFile)

	jsonObject := []models.Expenses{}
	if err = jsonParser.Decode(&jsonObject); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if jsonObject == nil {
		c.JSON(204, gin.H{"route": route, "method": method, "userid": userid, "user": ""})
		return
	}

	expenses := []models.Expenses{}

	for _, element := range jsonObject {
		if element.UserId == int32(userid) {
			expenses = append(expenses, element)
		}
	}

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
