package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"test-invoice/handler"
	"test-invoice/infrastructure"
	"test-invoice/middleware"
	"test-invoice/usecase"
	"test-invoice/usecase/queryservice"
)

func main() {
	infrastructure.Init()
	if err := godotenv.Load(".env"); err != nil {
		log.Panicf("load .env error: %v", err)
	}
	db := infrastructure.GetDB()

	r := gin.Default()

	users := r.Group("api/users")
	{
		userHandler := handler.NewUserUsecase(
			usecase.NewUserUsecase(infrastructure.NewUser(db)),
		)
		users.POST("/", userHandler.Create)
		users.POST("/login", userHandler.Login)
	}

	invoices := r.Group("api/invoices", middleware.JWTAuthMiddleware())
	{
		invoiceHandler := handler.NewInvoiceHandler(
			usecase.NewInvoiceUsecase(
				infrastructure.NewInvoice(db),
				infrastructure.NewUser(db),
				infrastructure.NewCustomer(db),
			),
			queryservice.NewInvoiceQuery(db),
		)
		invoices.GET("/", invoiceHandler.List)
		invoices.POST("/", invoiceHandler.Create)
	}

	r.Run()
}
