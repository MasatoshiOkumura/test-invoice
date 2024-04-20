package usecase

import (
	"net/http"
	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	"test-invoice/domain/service"
	errcode "test-invoice/lib"
)

type userUsecase struct {
	userRepo repository.User
}

type CreateUserInput struct {
	CompanyID int
	Name      string
	Mail      string
	Password  string
}

type LoginInput struct {
	Mail     string
	Password string
}

type UserUsecase interface {
	CreateUser(in *CreateUserInput) (*model.User, error)
	Login(in *LoginInput) (string, error)
}

func NewUserUsecase(userRepo repository.User) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) CreateUser(in *CreateUserInput) (*model.User, error) {
	b, err := service.IsExistMail(in.Mail)
	if b {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "mail is already exist")
	}
	if err != nil {
		return nil, err
	}

	user, err := model.NewUser(in.CompanyID, in.Name, in.Mail, in.Password)
	if err != nil {
		return nil, err
	}
	user, err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Login(in *LoginInput) (string, error) {
	user, err := u.userRepo.FindByMail(in.Mail)
	if err != nil {
		return "", err
	}

	token, err := u.userRepo.Login(in.Mail, user.Password, in.Password)
	if err != nil {
		return "", err
	}

	return token, nil
}
