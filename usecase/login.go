package usecase

import (
	"test-invoice/domain/repository"
)

type loginImpl struct {
	Mail           string
	Password       string
	UserRepository repository.User
}

type loginUsecase interface {
	Execute() (string, error)
}

func NewLoginUsecase(
	mail string,
	password string,
	repo repository.User,
) loginImpl {
	return loginImpl{
		Mail:           mail,
		Password:       password,
		UserRepository: repo,
	}
}

func (impl loginImpl) Execute() (string, error) {
	user, err := impl.UserRepository.FindByMail(impl.Mail)
	if err != nil {
		return "", err
	}

	token, err := impl.UserRepository.Login(impl.Mail, user.Password, impl.Password)
	if err != nil {
		return "", err
	}

	return token, nil
}
