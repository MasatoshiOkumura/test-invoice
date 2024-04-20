package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"test-invoice/infrastructure"
	"test-invoice/infrastructure/dto"
	errcode "test-invoice/lib"
	"test-invoice/usecase"
	"test-invoice/usecase/queryservice"
)

type InvoiceHandler struct{}

type InvoiceListInput struct {
	Date string
}

func (i InvoiceHandler) List(c *gin.Context) {
	db := infrastructure.GetDB()
	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid date query format. Please write as 2006-01-02",
		})
		return
	}
	if date.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "specify future date",
		})
		return
	}
	m, exist := c.Get("mail")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot get login user mail",
		})
	}
	mail, ok := m.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Mail is not string"})
		return
	}

	// 集約を跨って取得するためクエリサービスに実装
	in := &queryservice.ListInvoicesInput{
		Date: date,
	}
	invoices, err := queryservice.NewInvoiceQuery(db).List(mail, in)
	if err != nil {
		if e, ok := err.(*errcode.HTTPError); ok {
			c.JSON(e.Code, gin.H{"error": e.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, invoices)
}

// 金額データはstringをdecimalに変換する
type InvoiceCreateInput struct {
	CustomerID int    `json:"customer_id"`
	Payment    string `json:"payment"`
	FeeRate    string `json:"fee_rate"`
	Deadline   string `json:"deadline"`
}

func (h InvoiceHandler) Create(c *gin.Context) {
	var input InvoiceCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	in := &usecase.CreateInvoiceInput{
		CustomerID: input.CustomerID,
		Payment:    input.Payment,
		FeeRate:    input.FeeRate,
		Deadline:   input.Deadline,
	}
	m, exist := c.Get("mail")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot get login user mail",
		})
	}
	mail, ok := m.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Mail is not string"})
		return
	}

	i, err := usecase.NewInvoiceUsecase(
		infrastructure.NewInvoice(),
		infrastructure.NewUser(),
		infrastructure.NewCustomer(),
	).Create(mail, in)
	if err != nil {
		if e, ok := err.(*errcode.HTTPError); ok {
			c.JSON(e.Code, gin.H{"error": e.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	invoice := dto.ConvertToInvoiceDTO(i)

	c.JSON(http.StatusOK, invoice)
}
