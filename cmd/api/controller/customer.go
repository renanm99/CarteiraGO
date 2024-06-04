package controller

import (
	"carteirago/cmd/api/db/repository"
	"carteirago/cmd/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CustomerGET(c *gin.Context) {
	route := "customer"
	method := c.Request.Method
	id := c.DefaultQuery("id", "")
	userid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	code, customer, err := repository.CustomerGET(userid)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if customer.Id == 0 {
		c.JSON(204, gin.H{"route": route, "method": method, "userid": userid, "incomes": ""})
		return
	}

	if id == "" {
		c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "customer": customer})
		return
	}
}

func CustomerPOST(c *gin.Context) {
	route := "customer"
	method := c.Request.Method
	userid, err := strconv.Atoi(c.Query("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := new(models.Customer)

	err = c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer.Id = int32(userid)

	code, err := repository.CustomerPOST(customer)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "customer": customer})
}