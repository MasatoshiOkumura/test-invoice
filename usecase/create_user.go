package usecase

import (
	"net/http"

	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	"test-invoice/domain/service"
	errcode "test-invoice/lib"
)

type createUserUsecaseImpl struct {
	CompanyID      int
	Name           string
	Mail           string
	Password       string
	UserRepository repository.User
}

type createUserUsecase interface {
	Execute() (*model.User, error)
}

func NewCreateUserUsecase(
	companyID int,
	name string,
	mail string,
	password string,
	repo repository.User,
) createUserUsecase {
	return createUserUsecaseImpl{
		CompanyID:      companyID,
		Name:           name,
		Mail:           mail,
		Password:       password,
		UserRepository: repo,
	}
}

func (impl createUserUsecaseImpl) Execute() (*model.User, error) {
	b, err := service.IsExistMail(impl.Mail)
	if b {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "mail is already exist")
	}
	if err != nil {
		return nil, err
	}

	u, err := impl.UserRepository.Create(impl.CompanyID, impl.Name, impl.Mail, impl.Password)
	if err != nil {
		return nil, err
	}

	return u, nil
}
