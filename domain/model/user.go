package model

import "time"

type User struct {
	ID        int       `json:"id"`
	CompanyID int       `json:"company_id"`
	Name      string    `json:"name"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(companyID int, name string, mail string, password string) (*User, error) {
	user := &User{
		CompanyID: companyID,
		Name:      name,
		Mail:      mail,
		Password:  password,
	}

	return user, nil
}
