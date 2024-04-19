package model

import (
	"fmt"
	"time"
)

type AccountNumber string

type Bank struct {
	ID            int           `json:"id"`
	CustomerID    int           `json:"customer_id"`
	Name          string        `json:"name"`
	AccountNumber AccountNumber `json:"account_number"`
	AccountName   string        `json:"account_name"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

func NewAccountNumber(n string) (AccountNumber, error) {
	if len(n) > 20 {
		return "", fmt.Errorf("account_number must be less than 20: %v", n)
	}
	return AccountNumber(n), nil
}
