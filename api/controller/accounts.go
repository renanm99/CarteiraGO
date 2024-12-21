package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renanm99/carteirago/api/db"
	"github.com/renanm99/carteirago/api/db/repository"
	"github.com/renanm99/carteirago/api/models"
)

func SlashGet(c *gin.Context) {
	_, err := db.Database()
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	c.String(http.StatusOK, "Hi")
	c.Done()
}

func AccountsGET(c *gin.Context) {
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

	accountType := c.Request.Header["Account"][0]
	if accountType == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no header account type"})
		return
	}

	code, accounts, err := repository.AccountsSelect(userid, accountType)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if len(accounts) == 0 {
		c.JSON(http.StatusOK, gin.H{"accounts": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
	c.Done()
}

func AccountsPOST(c *gin.Context) {
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

	account := new(models.Accounts)
	err = c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account.UserId = userid

	code, err := repository.AccountsInsert(account)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "")
	c.Done()

}

func AccountsPUT(c *gin.Context) {
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

	account := new(models.Accounts)

	err = c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account.UserId = userid

	code, err := repository.AccountsUpdate(account)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, "")
	c.Done()

}

func AccountsDELETE(c *gin.Context) {

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

	account := new(models.Accounts)
	err = c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accountType := c.Request.Header["Account"][0]
	if accountType == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no header account type"})
		return
	}

	code, err := repository.AccountsDelete(userid, account.Id, accountType)

	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
	c.Done()

}

func DashboardAccounts(c *gin.Context) {
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

	account := c.Request.Header["Account"][0]
	if account == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no header account type"})
		return
	}

	code, accounts, err := repository.DashboardSelect(userid, account)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}

	if len(accounts) == 0 {
		c.JSON(http.StatusOK, gin.H{"dashboard": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dashboard": accounts})
	c.Done()
}
