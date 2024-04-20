package repository

import "test-invoice/domain/model"

type Company interface {
	FindByID(id int) (*model.Company, error)
}
