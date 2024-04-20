package repository

import "test-invoice/domain/model"

type Invoice interface {
	Create(*model.Invoice) (*model.Invoice, error)
}
