package main

import (
	"carteirago/cmd/api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/expenses", controller.ExpensesGET)
	r.POST("/expenses", controller.ExpensesPOST)

	r.GET("/incomes", controller.IncomesGET)
	r.POST("/incomes", controller.IncomesPOST)

	r.GET("/user", controller.UserGET)
	r.POST("/user", controller.UserPOST)

	r.Run(":8080")
}
