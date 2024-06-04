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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, customer, err := repository.CustomerSelect(userid)
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

	customer.Id = userid

	code, err := repository.CustomerInsert(customer)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "customer": customer})
}

func CustomerPUT(c *gin.Context) {
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

	customer.Id = userid

	code, err := repository.CustomerUpdate(customer)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"route": route, "method": method, "userid": userid, "customer": customer})
}

func CustomerDELETE(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("userid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, err := repository.CustomerDelete(userid)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
	c.Done()
}
