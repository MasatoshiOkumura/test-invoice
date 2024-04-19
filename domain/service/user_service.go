package service

import (
	"test-invoice/domain/model"
	"test-invoice/infrastructure"
)

func IsExistMail(mail string) (bool, error) {
	db := infrastructure.GetDB()

	var users []*model.User
	if err := db.Find(&users, "mail = ?", mail).Error; err != nil {
		return false, err
	}

	if len(users) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
