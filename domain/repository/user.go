package repository

import "test-invoice/domain/model"

type User interface {
	Create(*model.User) (*model.User, error)
	Login(mail string, password string, inPassword string) (string, error)
	FindByMail(mail string) (*model.User, error)
}
