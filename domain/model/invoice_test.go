package model_test

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"test-invoice/domain/model"
)

func TestInvoice_NewInvoice(t *testing.T) {
	type args struct {
		companyID  int
		customerID int
		payment    string
		feeRate    string
		deadline   string
	}
	tests := []struct {
		name       string
		args       args
		want       func() model.Invoice
		wantErrMsg string
	}{
		{
			name: "正常系",
			args: args{
				companyID:  1,
				customerID: 1,
				payment:    "10000",
				feeRate:    "0.04",
				deadline:   time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			},
			want: func() model.Invoice {
				payment, _ := decimal.NewFromString("10000")
				fee, _ := decimal.NewFromString("400")
				tax, _ := decimal.NewFromString("1040")
				billingAmount, _ := decimal.NewFromString("10440")
				return model.Invoice{
					CompanyID:     1,
					CustomerID:    1,
					Payment:       model.InvoicePayment(payment),
					Fee:           model.InvoiceFee(fee),
					Tax:           model.InvoiceTax(tax),
					BillingAmount: model.InvoiceBillingAmount(billingAmount),
					Status:        1,
					Deadline:      time.Now().AddDate(0, 0, 1),
				}
			},
		},
		{
			name: "支払金額に小数点が含まれる場合",
			args: args{
				companyID:  1,
				customerID: 1,
				payment:    "10000.55",
				feeRate:    "0.04",
				deadline:   time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			},
			want:       nil,
			wantErrMsg: "payment format is not correct",
		},
		{
			name: "支払金額の桁が大きすぎる場合エラーを返す",
			args: args{
				companyID:  1,
				customerID: 1,
				payment:    "1234567901234567",
				feeRate:    "0.04",
				deadline:   time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			},
			want:       nil,
			wantErrMsg: "payment format is not correct",
		},
		{
			name: "手数料率の桁が大きすぎる場合エラーを返す",
			args: args{
				companyID:  1,
				customerID: 1,
				payment:    "10000",
				feeRate:    "10",
				deadline:   time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			},
			want:       nil,
			wantErrMsg: "fee_rate format is not correct",
		},
		{
			name: "手数料率が小数点第3位まで指定されている場合エラーを返す",
			args: args{
				companyID:  1,
				customerID: 1,
				payment:    "10000",
				feeRate:    "0.111",
				deadline:   time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
			},
			want:       nil,
			wantErrMsg: "fee_rate format is not correct",
		},
		{
			name: "過去の支払期日を指定した場合エラーを返す",
			args: args{
				companyID:  1,
				customerID: 1,
				payment:    "10000",
				feeRate:    "0.04",
				deadline:   time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
			},
			want:       nil,
			wantErrMsg: "deadline must be after now",
		},
		{
			name: "支払期日がフォーマット外の場合エラーを返す",
			args: args{
				companyID:  1,
				customerID: 1,
				payment:    "10000",
				feeRate:    "0.04",
				deadline:   "aaa",
			},
			want:       nil,
			wantErrMsg: "invalid deadline format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			act, err := model.NewInvoice(tt.args.companyID, tt.args.customerID, tt.args.payment, tt.args.feeRate, tt.args.deadline)

			if len(tt.wantErrMsg) > 0 {
				assert.ErrorContains(t, err, tt.wantErrMsg)
				return
			}
			want := tt.want()
			assert.Equal(t, want.CompanyID, act.CompanyID)
			assert.Equal(t, want.CustomerID, act.CustomerID)
			assert.Equal(t, want.Payment, act.Payment)
			assert.Equal(t, want.Fee, act.Fee)
			assert.Equal(t, want.Tax, act.Tax)
			assert.Equal(t, want.BillingAmount, act.BillingAmount)
			assert.Equal(t, want.Deadline.Format("2006-01-02"), act.Deadline.Format("2006-01-02"))
		})
	}
}
