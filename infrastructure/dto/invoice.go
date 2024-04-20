package dto

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"

	"test-invoice/domain/model"
)

type Invoice struct {
	ID            int       `json:"id"`
	CompanyID     int       `json:"company_id"`
	CustomerID    int       `json:"customer_id"`
	IssueDate     time.Time `json:"issue_date"`
	Payment       string    `json:"payment"`
	Fee           string    `json:"fee"`
	FeeRate       string    `json:"fee_rate"`
	Tax           string    `json:"tax"`
	TaxRate       string    `json:"tax_rate"`
	BillingAmount string    `json:"billing_amount"`
	Deadline      time.Time `json:"deadline"`
	Status        int       `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type InvoiceDTO struct {
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
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (i *Invoice) ConvertToModel() (*model.Invoice, error) {
	p, err := decimal.NewFromString(i.Payment)
	if err != nil {
		return nil, err
	}
	f, err := decimal.NewFromString(i.Fee)
	if err != nil {
		return nil, err
	}
	fr, err := decimal.NewFromString(i.FeeRate)
	if err != nil {
		return nil, err
	}
	t, err := decimal.NewFromString(i.Tax)
	if err != nil {
		return nil, err
	}
	tr, err := decimal.NewFromString(i.TaxRate)
	if err != nil {
		return nil, err
	}
	ba, err := decimal.NewFromString(i.BillingAmount)
	if err != nil {
		return nil, err
	}

	status := model.NewInvoiceStatus(i.Status)
	if status == model.InvoiceStatusUnknown {
		return nil, errors.New("invalid invoice status")
	}

	return &model.Invoice{
		ID:            i.ID,
		CompanyID:     i.CompanyID,
		CustomerID:    i.CustomerID,
		IssueDate:     i.IssueDate,
		Payment:       model.InvoicePayment(p),
		Fee:           model.InvoiceFee(f),
		FeeRate:       model.InvoiceFeeRate(fr),
		Tax:           model.InvoiceTax(t),
		TaxRate:       model.InvoiceTaxRate(tr),
		BillingAmount: model.InvoiceBillingAmount(ba),
		Deadline:      i.Deadline,
		Status:        status,
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}, nil
}

func ConvertToInvoiceDTO(i *model.Invoice) InvoiceDTO {
	return InvoiceDTO{
		ID:            i.ID,
		CompanyID:     i.CompanyID,
		CustomerID:    i.CustomerID,
		IssueDate:     i.IssueDate.Format("2006-01-02"),
		Payment:       decimal.Decimal(i.Payment).String(),
		Fee:           decimal.Decimal(i.Fee).String(),
		FeeRate:       decimal.Decimal(i.FeeRate).String(),
		Tax:           decimal.Decimal(i.Tax).String(),
		TaxRate:       decimal.Decimal(i.TaxRate).String(),
		BillingAmount: decimal.Decimal(i.BillingAmount).String(),
		Deadline:      i.IssueDate.Format("2006-01-02"),
		Status:        int(i.Status),
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}
}
