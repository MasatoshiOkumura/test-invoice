package infrastructure

import (
	"net/http"

	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	"test-invoice/infrastructure/dto"
	errcode "test-invoice/lib"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomer(db *gorm.DB) repository.Customer {
	return &customerRepo{db: db}
}

func (c *customerRepo) FindByID(id int) (*model.Customer, error) {
	customerDAO := dto.Customer{}
	bankDAO := dto.Bank{}

	if err := c.db.First(&customerDAO, id).Error; err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "bank is not exist")
	}
	// BankはCustomer集約のため、必ずCustomerに含めた状態で返す
	if err := c.db.Where("customer_id = ?", customerDAO.ID).First(&bankDAO).Error; err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "bank is not exist")
	}

	customer, err := customerDAO.ConvertToModel(&bankDAO)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
