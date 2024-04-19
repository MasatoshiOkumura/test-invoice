package infrastructure

import (
	"golang.org/x/crypto/bcrypt"

	"test-invoice/domain/model"
	"test-invoice/infrastructure/dto"
	"test-invoice/repository"
)

type user struct{}

func NewUser() repository.User {
	return &user{}
}

func (user *user) Create(companyID int, name string, mail string, password string) (*model.User, error) {
	db := GetDB()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userDAO := dto.User{
		CompanyID: companyID,
		Name:      name,
		Mail:      mail,
		Password:  string(hash),
	}

	if err := db.Create(&userDAO).Error; err != nil {
		return nil, err
	}

	return userDAO.ConvertToModel(), nil
}
