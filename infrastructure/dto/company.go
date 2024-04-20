package dto

import (
	"test-invoice/domain/model"
	"time"
)

type Company struct {
	ID            int       `json:"id" binding:"required"`
	Name          string    `json:"name"`
	Repesentative string    `json:"repesentative"`
	Tel           string    `json:"tel"`
	PostCode      string    `json:"post_code"`
	Address       string    `json:"address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (c *Company) ConvertToModel() *model.Company {
	return &model.Company{
		ID:            c.ID,
		Name:          c.Name,
		Repesentative: c.Repesentative,
		Tel:           c.Tel,
		PostCode:      c.PostCode,
		Address:       c.Address,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
	}
}
