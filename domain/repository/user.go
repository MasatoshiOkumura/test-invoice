package repository

import "test-invoice/domain/model"

type User interface {
	Create(companyID int, name string, mail string, password string) (*model.User, error)
	Login(mail string, password string, inPassword string) (string, error)
	FindByMail(mail string) (*model.User, error)
}
