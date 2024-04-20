package model

import "time"

type Customer struct {
	ID            int    `json:"id"`
	CompanyID     int    `json:"company_id"`
	Name          string `json:"name"`
	Repesentative string `json:"repesentative"`
	Tel           string `json:"tel"`
	PostCode      string `json:"post_code"`
	Address       string `json:"address"`
	// BankはCustomerに対してのみ意味を持つ(独立してアクセスされないため、Customerに含める)
	Bank      *Bank
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
