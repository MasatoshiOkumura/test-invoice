package dto

import (
	"time"

	"test-invoice/domain/model"
)

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

func (c *Customer) ConvertToModel(b *Bank) (*model.Customer, error) {
	bank := &model.Bank{
		ID:            b.ID,
		CustomerID:    b.CustomerID,
		Name:          b.Name,
		AccountNumber: model.AccountNumber(b.AccountNumber),
		AccountName:   b.AccountName,
		CreatedAt:     b.CreatedAt,
		UpdatedAt:     b.UpdatedAt,
	}

	return &model.Customer{
		ID:            c.ID,
		CompanyID:     c.CompanyID,
		Name:          c.Name,
		Repesentative: c.Repesentative,
		Tel:           c.Tel,
		PostCode:      c.PostCode,
		Address:       c.Address,
		Bank:          bank,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}, nil
}
