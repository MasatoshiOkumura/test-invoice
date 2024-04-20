package dto

// 金額データはstringをdecimalに変換する
type InvoiceCreateInput struct {
	CompanyID  int
	CustomerID int
	Payment    string
	FeeRate    string
	Deadline   string
}
