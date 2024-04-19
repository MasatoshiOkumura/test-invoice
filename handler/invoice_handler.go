package handler

import "github.com/gin-gonic/gin"

type InvoiceHandler struct{}

func (i InvoiceHandler) List(c *gin.Context) {
	// TODO
	c.JSON(200, gin.H{
		"message": "test list",
	})
}

func (i InvoiceHandler) Create(c *gin.Context) {
	// TODO
	c.JSON(200, gin.H{
		"message": "test create",
	})
}
