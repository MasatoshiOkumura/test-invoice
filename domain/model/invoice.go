package model

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/shopspring/decimal"

	errcode "test-invoice/lib"
)

type InvoicePayment decimal.Decimal
type InvoiceFee decimal.Decimal
type InvoiceFeeRate decimal.Decimal
type InvoiceTax decimal.Decimal
type InvoiceTaxRate decimal.Decimal
type InvoiceBillingAmount decimal.Decimal
type InvoiceStatus int

const (
	InvoiceStatusUnknown    InvoiceStatus = 0
	InvoiceStausUnprocessed InvoiceStatus = 1
	InvoiceStatusInprogress InvoiceStatus = 2
	InvoiceStatusDone       InvoiceStatus = 3
	IncoiceStatusError      InvoiceStatus = 4
)

var TaxRate = decimal.NewFromFloat(0.1)

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
	d, err := validateDecimalPrecision(value, 15, 0)
	if err != nil {
		return InvoicePayment(decimal.NewFromInt(0)), err
	}
	return InvoicePayment(d), nil
}

func NewInvoiceFee(value string) (InvoiceFee, error) {
	d, err := validateDecimalPrecision(value, 12, 0)
	if err != nil {
		return InvoiceFee(decimal.NewFromInt(0)), err
	}
	return InvoiceFee(d), nil
}

func NewInvoiceFeeRate(value string) (InvoiceFeeRate, error) {
	d, err := validateDecimalPrecision(value, 3, 2)
	if err != nil {
		return InvoiceFeeRate(decimal.NewFromInt(0)), err
	}
	if d.GreaterThan(decimal.NewFromInt(1)) {
		return InvoiceFeeRate(decimal.NewFromInt(0)), errors.New("fee_rate must be less than 1")
	}
	return InvoiceFeeRate(d), nil
}

func NewInvoiceTax(value string) (InvoiceTax, error) {
	d, err := validateDecimalPrecision(value, 12, 0)
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
	d, err := validateDecimalPrecision(value, 15, 0)
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

func newDeadline(deadline string) (*time.Time, error) {
	in, err := time.Parse("2006-01-02", deadline)
	if err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "invalid deadline format")
	}
	if time.Now().After(in) {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "deadline must be after now")
	}
	return &in, nil
}

func NewInvoice(companyID int, customerID int, payment string, feeRate string, deadline string) (*Invoice, error) {
	p, err := NewInvoicePayment(payment)
	if err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "payment format is not correct")
	}
	fr, err := NewInvoiceFeeRate(feeRate)
	if err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "fee_rate format is not correct")
	}
	dl, err := newDeadline(deadline)
	if err != nil {
		return nil, err
	}

	pDeci := decimal.Decimal(p)
	frDeci := decimal.Decimal(fr)
	// 請求金額=請求金額+(請求金額*手数料率*消費税率)
	ba := pDeci.Add(pDeci.Mul(frDeci).Mul(TaxRate.Add(decimal.NewFromInt(1))))
	ba = ba.Ceil()
	// 手数料=請求金額*手数料率
	fDeci := pDeci.Mul(frDeci)
	fDeci = fDeci.Ceil()
	// 消費税=(請求金額+手数料)*消費税率
	tax := (pDeci.Add(fDeci)).Mul(TaxRate)
	tax = tax.Ceil()

	today, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		return nil, errcode.NewHTTPError(http.StatusInternalServerError, "time now parse failed")
	}

	invoice := &Invoice{
		CompanyID:     companyID,
		CustomerID:    customerID,
		IssueDate:     today,
		Payment:       p,
		Fee:           InvoiceFee(fDeci),
		FeeRate:       fr,
		Tax:           InvoiceTax(tax),
		TaxRate:       InvoiceTaxRate(TaxRate),
		BillingAmount: InvoiceBillingAmount(ba),
		Deadline:      *dl,
		Status:        InvoiceStausUnprocessed,
	}

	return invoice, nil
}

func NewInvoiceStatus(n int) InvoiceStatus {
	switch n {
	case int(InvoiceStausUnprocessed):
		return InvoiceStausUnprocessed
	case int(InvoiceStatusInprogress):
		return InvoiceStatusInprogress
	case int(InvoiceStatusDone):
		return InvoiceStatusDone
	default:
		return InvoiceStatusUnknown
	}
}
