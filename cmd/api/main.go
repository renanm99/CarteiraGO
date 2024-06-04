package main

import (
	"carteirago/cmd/api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

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

	r.Run(":8080")
}
