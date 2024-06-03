package repository

import (
	"carteirago/cmd/api/models"
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserGET(c *gin.Context) {
	route := "user"
	method := c.Request.Method
	id := c.DefaultQuery("id", "")
	userid, err := strconv.ParseInt(c.Query("userid"), 10, 8)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonFile, err := os.Open("/users.json")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	jsonParser := json.NewDecoder(jsonFile)

	jsonObject := []models.User{}
	if err = jsonParser.Decode(&jsonObject); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if jsonObject == nil {
		c.JSON(204, gin.H{"route": route, "method": method, "userid": userid, "user": ""})
		return
	}

	user := models.User{}

	for _, element := range jsonObject {
		if element.Id == int32(userid) {
			user = element
		}
	}

	if user.Id == 0 {
		c.JSON(204, gin.H{"route": route, "method": method, "userid": userid, "incomes": ""})
		return
	}

	if id == "" {
		c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "user": user})
		return
	}
}

func UserPOST(c *gin.Context) {
	route := "user"
	method := c.Request.Method
	userid := c.Query("userid")
	id := c.Query("id")
	c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "id": id})
}
