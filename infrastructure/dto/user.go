package dto

import (
	"test-invoice/domain/model"
	"time"
)

type UserCreateInput struct {
	CompanyID int    `json:"company_id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
}

type User struct {
	ID        int `json:"id" binding:"required"`
	CompanyID int
	Name      string
	Mail      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDTO struct {
	ID        int       `json:"id"`
	CompanyID int       `json:"company_id"`
	Name      string    `json:"name"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ConvertToModel() *model.User {
	return &model.User{
		ID:        u.ID,
		CompanyID: u.CompanyID,
		Name:      u.Name,
		Mail:      u.Mail,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ConvertToUserDTO(u *model.User) UserDTO {
	return UserDTO{
		ID:        u.ID,
		CompanyID: u.CompanyID,
		Name:      u.Name,
		Mail:      u.Mail,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
