package queryservice

import (
	"net/http"
	"time"

	"gorm.io/gorm"

	"test-invoice/domain/model"
	errcode "test-invoice/lib"
)

type ListInvoicesInput struct {
	Date time.Time
}

// 参照用に独自の型を使用
type Invoice struct {
	ID            int       `json:"id"`
	CompanyID     int       `json:"company_id"`
	CustomerID    int       `json:"customer_id"`
	IssueDate     string    `json:"issue_date"`
	Payment       string    `json:"payment"`
	Fee           string    `json:"fee"`
	FeeRate       string    `json:"fee_rate"`
	Tax           string    `json:"tax"`
	TaxRate       string    `json:"tax_rate"`
	BillingAmount string    `json:"billing_amount"`
	Deadline      string    `json:"deadline"`
	Status        int       `json:"status"`
	Customer      *Customer `json:"customer"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type Customer struct {
	ID            int       `json:"id"`
	CompanyID     int       `json:"company_id"`
	Name          string    `json:"name"`
	Repesentative string    `json:"repesentative"`
	Tel           string    `json:"tel"`
	PostCode      string    `json:"post_code"`
	Address       string    `json:"address"`
	Bank          *Bank     `json:"bank"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type Bank struct {
	ID            int       `json:"id"`
	CustomerID    int       `json:"customer_id"`
	Name          string    `json:"name"`
	AccountNumber string    `json:"account_number"`
	AccountName   string    `json:"account_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type invoiceQuery struct {
	db *gorm.DB
}

type InvoiceQuery interface {
	List(mail string, in *ListInvoicesInput) ([]*Invoice, error)
}

func NewInvoiceQuery(db *gorm.DB) InvoiceQuery {
	return &invoiceQuery{db: db}
}

func (q *invoiceQuery) List(mail string, in *ListInvoicesInput) ([]*Invoice, error) {
	// ユーザーの属する会社を取得
	user := &model.User{}
	if err := q.db.Where("mail = ?", mail).First(&user).Error; err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "user is not exist")
	}

	// 期間内かつ未払いの請求書を取得
	invoices := []*Invoice{}
	if err := q.db.Preload("Customer.Bank").
		Where("company_id = ? AND deadline < ? AND status = ?", user.CompanyID, in.Date, model.InvoiceStausUnprocessed).
		Find(&invoices).Error; err != nil {
		return nil, errcode.NewHTTPError(http.StatusInternalServerError, "error find invoices")
	}
	if len(invoices) == 0 {
		return nil, errcode.NewHTTPError(http.StatusNotFound, "invoices not found")
	}

	return invoices, nil
}
