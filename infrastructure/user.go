package infrastructure

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	"test-invoice/infrastructure/dto"
	errcode "test-invoice/lib"
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

func (u *user) Login(mail string, password string, inPassword string) (string, error) {
	key := os.Getenv("ACCESS_SECRET_KEY")

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(inPassword)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mail": mail,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	accessToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (u *user) FindByMail(mail string) (*model.User, error) {
	db := GetDB()
	userDAO := dto.User{}

	if err := db.Where("mail = ?", mail).First(&userDAO).Error; err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "user is not exist")
	}

	return userDAO.ConvertToModel(), nil
}
