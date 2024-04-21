package infrastructure

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	"test-invoice/infrastructure/dto"
	errcode "test-invoice/lib"
)

type userRepo struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) repository.User {
	return &userRepo{db: db}
}

func (u *userRepo) Create(user *model.User) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userDAO := dto.User{
		CompanyID: user.CompanyID,
		Name:      user.Name,
		Mail:      user.Mail,
		Password:  string(hash),
	}

	if err := u.db.Create(&userDAO).Error; err != nil {
		return nil, err
	}

	return userDAO.ConvertToModel(), nil
}

func (userRepo *userRepo) Login(mail string, password string, inPassword string) (string, error) {
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

func (u *userRepo) FindByMail(mail string) (*model.User, error) {
	db := GetDB()
	userDAO := dto.User{}

	if err := db.Where("mail = ?", mail).First(&userDAO).Error; err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "user is not exist")
	}

	return userDAO.ConvertToModel(), nil
}
