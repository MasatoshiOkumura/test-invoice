package repository

import "test-invoice/domain/model"

type Customer interface {
	FindByID(id int) (*model.Customer, error)
}
