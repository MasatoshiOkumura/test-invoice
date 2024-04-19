package repository

import "test-invoice/domain/model"

type User interface {
	Create(companyID int, name string, mail string, password string) (*model.User, error)
}
