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

func (h UserHandler) Create(c *gin.Context) {
	var input dto.UserCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := infrastructure.NewUser()
	u, err := usecase.NewCreateUserUsecase(input.CompanyID, input.Name, input.Mail, input.Password, repo).Execute()
	if err != nil {
		if e, ok := err.(*errcode.HTTPError); ok {
			c.JSON(e.Code, gin.H{"error": e.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
	}

	user := dto.ConvertToUserDTO(u)

	c.JSON(http.StatusOK, user)
}

func (h UserHandler) Login(c *gin.Context) {
	var input dto.UserLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := infrastructure.NewUser()
	token, err := usecase.NewLoginUsecase(input.Mail, input.Password, repo).Execute()
	if err != nil {
		if e, ok := err.(*errcode.HTTPError); ok {
			c.JSON(e.Code, gin.H{"error": e.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
