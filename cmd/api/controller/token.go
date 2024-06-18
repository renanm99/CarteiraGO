package controller

import (
	"carteirago/cmd/api/db/repository"
	"carteirago/cmd/api/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func LoginHandler(c *gin.Context) {

	user := new(models.Customer)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := repository.CheckPwd(user.Email, user.Password)
	if id > 0 {
		tokenString, err := CreateToken(user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		setCookieHandler(c, strconv.Itoa(id), tokenString)
		c.JSON(http.StatusOK, gin.H{"code": 200})
		c.Done()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
}

func DeleteCookie(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "localhost", false, true)
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.String(http.StatusNoContent, "log out")
	c.Done()
	//c.String(http.StatusOK, "Cookie has been deleted")
}

func setCookieHandler(c *gin.Context, email string, jwt string) {
	c.SetSameSite(http.SameSiteStrictMode)
	_, _, err := getCookieHandler(c)
	if err != nil {
		c.SetCookie("user", email, 3600, "/", "localhost", false, true)
		c.SetCookie("token", jwt, 3600, "/", "localhost", false, true)
	}
}

func getCookieHandler(c *gin.Context) (string, string, error) {
	user, err := c.Cookie("user")
	if err != nil {
		//c.String(http.StatusNotFound, "not found")
		return "", "", err
	}
	cookietoken, err := c.Cookie("token")
	if err != nil {
		//c.String(http.StatusNotFound, "not found")
		return "", "", err
	}
	return user, cookietoken, nil
	//c.String(http.StatusOK, "Cookie value: %s", cookietoken)
}

func CheckSignin(c *gin.Context) {
	_, _, err := getCookieHandler(c)
	if err != nil {
		c.String(http.StatusNotAcceptable, "not auth")
		c.Done()
		return
	}
	c.String(http.StatusOK, "log in")
	c.Done()
}
