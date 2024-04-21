package infrastructure

import (
	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	"test-invoice/infrastructure/dto"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type invoiceRepo struct {
	db *gorm.DB
}

func NewInvoice(db *gorm.DB) repository.Invoice {
	return &invoiceRepo{db: db}
}

func (i *invoiceRepo) Create(invoice *model.Invoice) (*model.Invoice, error) {
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

	if err := i.db.Create(&invoiceDAO).Error; err != nil {
		return nil, err
	}

	invoice, err := invoiceDAO.ConvertToModel()
	if err != nil {
		return nil, err
	}
	return invoice, nil
}
