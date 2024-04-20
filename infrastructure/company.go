package infrastructure

import (
	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	"test-invoice/infrastructure/dto"
)

type companyRepo struct{}

func NewCompany() repository.Company {
	return &companyRepo{}
}

func (c *companyRepo) FindByID(id int) (*model.Company, error) {
	db := GetDB()
	companyDAO := dto.Company{}

	if err := db.First(&companyDAO, id).Error; err != nil {
		return nil, err
	}

	company := companyDAO.ConvertToModel()
	return company, nil
}
