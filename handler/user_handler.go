package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"test-invoice/infrastructure/dto"
	errcode "test-invoice/lib"
	"test-invoice/usecase"
)

type UserHandler interface {
	Login(c *gin.Context)
	Create(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserUsecase(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

type UserCreateInput struct {
	CompanyID int    `json:"company_id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
}

func (h userHandler) Create(c *gin.Context) {
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
	u, err := h.userUsecase.CreateUser(in)
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

func (h userHandler) Login(c *gin.Context) {
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
	token, err := h.userUsecase.Login(in)
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
