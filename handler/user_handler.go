package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"test-invoice/infrastructure"
	"test-invoice/infrastructure/dto"
	errcode "test-invoice/lib"
	"test-invoice/usecase"
)

type UserHandler struct{}

type UserCreateInput struct {
	CompanyID int    `json:"company_id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
}

func (h UserHandler) Create(c *gin.Context) {
	var input UserCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	in := &usecase.CreateUserInput{
		CompanyID: input.CompanyID,
		Name:      input.Name,
		Mail:      input.Mail,
		Password:  input.Password,
	}
	repo := infrastructure.NewUser()
	u, err := usecase.NewUserUsecase(repo).CreateUser(in)
	if err != nil {
		if e, ok := err.(*errcode.HTTPError); ok {
			c.JSON(e.Code, gin.H{"error": e.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	user := dto.ConvertToUserDTO(u)

	c.JSON(http.StatusOK, user)
}

type UserLoginInput struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (h UserHandler) Login(c *gin.Context) {
	var input UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	in := &usecase.LoginInput{
		Mail:     input.Mail,
		Password: input.Password,
	}
	repo := infrastructure.NewUser()
	token, err := usecase.NewUserUsecase(repo).Login(in)
	if err != nil {
		if e, ok := err.(*errcode.HTTPError); ok {
			c.JSON(e.Code, gin.H{"error": e.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
