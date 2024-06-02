package main

import (
	"carteirago/cmd/api/controller"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	time := time.Now()
	fmt.Println(time)

	r := gin.Default()

	r.GET("/expenses", controller.ExpensesGET)
	r.POST("/expenses", controller.ExpensesPOST)

	r.GET("/incomes", controller.IncomesGET)
	r.POST("/incomes", controller.IncomesPOST)

	r.GET("/user", controller.UserGET)
	r.POST("/user", controller.UserPOST)

	r.Run(":8080")
}
