package model

import "time"

type Company struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Repesentative string    `json:"repesentative"`
	Tel           string    `json:"tel"`
	PostCode      string    `json:"post_code"`
	Address       string    `json:"address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
