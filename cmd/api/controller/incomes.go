package controller

import (
	"carteirago/cmd/api/models"
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const repo string = "C:\\dev\\carteirago\\cmd\\api\\db\\repo_data"

func IncomesGET(c *gin.Context) {
	route := "incomes"
	method := c.Request.Method
	id := c.DefaultQuery("id", "")
	userid, err := strconv.ParseInt(c.Query("userid"), 10, 8)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonFile, err := os.Open(repo + "/incomes.json")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonParser := json.NewDecoder(jsonFile)

	jsonObject := []models.Incomes{}
	if err = jsonParser.Decode(&jsonObject); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if len(jsonObject) == 0 {
		c.JSON(204, gin.H{"route": route, "method": method, "userid": userid, "incomes": nil})
		return
	}

	incomes := []models.Incomes{}

	for _, element := range jsonObject {
		if element.UserId == int32(userid) {
			incomes = append(incomes, element)
		}
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
	userid := c.Query("userid")
	id := c.Query("id")
	c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "id": id})
}
