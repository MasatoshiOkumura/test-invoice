package main

import (
	"github.com/gin-gonic/gin"

	"test-invoice/handler"
	"test-invoice/infrastructure"
)

func main() {
	infrastructure.Init()

	r := gin.Default()

	users := r.Group("api/users")
	{
		userHandler := handler.UserHandler{}
		users.POST("/", userHandler.Create)
	}

	invoices := r.Group("api/invoices")
	{
		invoiceHandler := handler.InvoiceHandler{}
		invoices.GET("/", invoiceHandler.List)
		invoices.POST("/", invoiceHandler.Create)
	}

	r.Run()
}
