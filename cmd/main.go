package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"test-invoice/handler"
	"test-invoice/infrastructure"
	"test-invoice/middleware"
)

func main() {
	infrastructure.Init()
	if err := godotenv.Load(".env"); err != nil {
		log.Panicf("load .env error: %v", err)
	}

	r := gin.Default()

	users := r.Group("api/users")
	{
		userHandler := handler.UserHandler{}
		users.POST("/", userHandler.Create)
		users.POST("/login", userHandler.Login)
	}

	invoices := r.Group("api/invoices", middleware.JWTAuthMiddleware())
	{
		invoiceHandler := handler.InvoiceHandler{}
		invoices.GET("/", invoiceHandler.List)
		invoices.POST("/", invoiceHandler.Create)
	}

	r.Run()
}
