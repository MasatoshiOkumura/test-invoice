package main

import (
	"test-invoice/handler"
	"test-invoice/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.Init()

	r := gin.Default()

	invoices := r.Group("api/invoices")
	{
		invoiceHandler := handler.InvoiceHandler{}
		invoices.GET("/", invoiceHandler.List)
		invoices.POST("/", invoiceHandler.Create)
	}

	r.Run()
}
