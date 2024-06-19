package main

import (
	"carteirago/cmd/api/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Account", "Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           12 * time.Hour,
	}))

	/*
		r.GET("/expenses", controller.ExpensesGET)
		r.POST("/expenses", controller.ExpensesPOST)
		r.PUT("/expenses", controller.ExpensesPUT)
		r.DELETE("/expenses", controller.ExpensesDELETE)
	*/

	r.GET("/expenses", controller.AccountsGET)
	r.POST("/expenses", controller.AccountsPOST)
	r.PUT("/expenses", controller.AccountsPUT)
	r.DELETE("/expenses", controller.AccountsDELETE)

	r.GET("/incomes", controller.AccountsGET)
	r.POST("/incomes", controller.AccountsPOST)
	r.PUT("/incomes", controller.AccountsPUT)
	r.DELETE("/incomes", controller.AccountsDELETE)

	r.GET("/user", controller.CustomerGET)
	r.POST("/user", controller.CustomerPOST)
	r.PUT("/user", controller.CustomerPUT)
	r.DELETE("/user", controller.CustomerDELETE)

	r.GET("/signin", controller.CheckSignin)

	r.POST("/login", controller.LoginHandler)
	r.DELETE("/login", controller.DeleteCookie)

	r.GET("/Dash", controller.DashboardAccounts)

	r.Run(":8080")
}
