package model

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type InvoicePayment decimal.Decimal
type InvoiceFee decimal.Decimal
type InvoiceFeeRate decimal.Decimal
type InvoiceTax decimal.Decimal
type InvoiceTaxRate decimal.Decimal
type InvoiceBillingAmount decimal.Decimal
type InvoiceStatus int

const (
	InvoiceStausUnprocessed InvoiceStatus = 1
	InvoiceStatusInprogress InvoiceStatus = 2
	InvoiceStatusDone       InvoiceStatus = 3
	IncoiceStatusError      InvoiceStatus = 4
)

type Invoice struct {
	ID            int                  `json:"id"`
	CompanyID     int                  `json:"company_id"`
	CustomerID    int                  `json:"customer_id"`
	IssueDate     time.Time            `json:"issue_date"`
	Payment       InvoicePayment       `json:"payment"`
	Fee           InvoiceFee           `json:"fee"`
	FeeRate       InvoiceFeeRate       `json:"fee_rate"`
	Tax           InvoiceTax           `json:"tax"`
	TaxRate       InvoiceTaxRate       `json:"tax_rate"`
	BillingAmount InvoiceBillingAmount `json:"billing_amount"`
	Deadline      time.Time            `json:"deadline"`
	Status        InvoiceStatus        `json:"status"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

func NewInvoicePayment(value string) (InvoicePayment, error) {
	d, err := validateDecimalPrecision(value, 15, 2)
	if err != nil {
		return InvoicePayment(decimal.NewFromInt(0)), err
	}
	return InvoicePayment(d), nil
}

func NewInvoiceFee(value string) (InvoiceFee, error) {
	d, err := validateDecimalPrecision(value, 12, 2)
	if err != nil {
		return InvoiceFee(decimal.NewFromInt(0)), err
	}
	return InvoiceFee(d), nil
}

func NewInvoiceFeeRate(value string) (InvoiceFeeRate, error) {
	d, err := validateDecimalPrecision(value, 5, 2)
	if err != nil {
		return InvoiceFeeRate(decimal.NewFromInt(0)), err
	}
	return InvoiceFeeRate(d), nil
}

func NewInvoiceTax(value string) (InvoiceTax, error) {
	d, err := validateDecimalPrecision(value, 12, 2)
	if err != nil {
		return InvoiceTax(decimal.NewFromInt(0)), err
	}
	return InvoiceTax(d), nil
}

func NewInvoiceTaxRate(value string) (InvoiceTaxRate, error) {
	d, err := validateDecimalPrecision(value, 5, 2)
	if err != nil {
		return InvoiceTaxRate(decimal.NewFromInt(0)), err
	}
	return InvoiceTaxRate(d), nil
}

func NewInvoiceBillingAmount(value string) (InvoiceBillingAmount, error) {
	d, err := validateDecimalPrecision(value, 15, 2)
	if err != nil {
		return InvoiceBillingAmount(decimal.NewFromInt(0)), err
	}
	return InvoiceBillingAmount(d), nil
}

func validateDecimalPrecision(value string, precision, scale int32) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(value)
	if err != nil {
		return decimal.NewFromInt(0), fmt.Errorf("invalid decimal format: %v", err)
	}

	// 小数点以下の桁数チェック
	digitsAfterDecimal := -d.Exponent()
	if digitsAfterDecimal > scale {
		return decimal.NewFromInt(0), fmt.Errorf("decimal exceeds scale: %d digits allowed, but got %d", scale, -digitsAfterDecimal)
	}

	// 総桁数チェック
	digitsTotal := int32(len(d.Coefficient().String()))
	if digitsTotal > precision {
		return decimal.NewFromInt(0), fmt.Errorf("decimal exceeds precision: %d total digits allowed, but got %d", precision, digitsTotal)
	}

	return d, nil
}
