package infrastructure

import (
	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	"test-invoice/infrastructure/dto"

	"github.com/shopspring/decimal"
)

type invoiceRepo struct{}

func NewInvoice() repository.Invoice {
	return &invoiceRepo{}
}

func (i *invoiceRepo) Create(invoice *model.Invoice) (*model.Invoice, error) {
	db := GetDB()

	invoiceDAO := dto.Invoice{
		CompanyID:     invoice.CompanyID,
		CustomerID:    invoice.CustomerID,
		IssueDate:     invoice.IssueDate,
		Payment:       decimal.Decimal(invoice.Payment).String(),
		Fee:           decimal.Decimal(invoice.Fee).String(),
		FeeRate:       decimal.Decimal(invoice.FeeRate).String(),
		Tax:           decimal.Decimal(invoice.Tax).String(),
		TaxRate:       decimal.Decimal(invoice.TaxRate).String(),
		BillingAmount: decimal.Decimal(invoice.BillingAmount).String(),
		Deadline:      invoice.Deadline,
		Status:        int(invoice.Status),
	}

	if err := db.Create(&invoiceDAO).Error; err != nil {
		return nil, err
	}

	invoice, err := invoiceDAO.ConvertToModel()
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
