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
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/expenses", controller.ExpensesGET)
	r.POST("/expenses", controller.ExpensesPOST)
	r.PUT("/expenses", controller.ExpensesPUT)
	r.DELETE("/expenses", controller.ExpensesDELETE)

	r.GET("/incomes", controller.IncomesGET)
	r.POST("/incomes", controller.IncomesPOST)
	r.PUT("/incomes", controller.IncomesPUT)
	r.DELETE("/incomes", controller.IncomesDELETE)

	r.GET("/user", controller.CustomerGET)
	r.POST("/user", controller.CustomerPOST)
	r.PUT("/user", controller.CustomerPUT)
	r.DELETE("/user", controller.CustomerDELETE)

	r.GET("/signin", controller.CheckSignin)

	r.POST("/login", controller.LoginHandler)
	r.DELETE("/login", controller.DeleteCookie)

	r.Run(":8080")
}
