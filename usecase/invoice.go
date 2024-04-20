package usecase

import (
	"net/http"

	"test-invoice/domain/model"
	"test-invoice/domain/repository"
	errcode "test-invoice/lib"
)

type invoiceUsecase struct {
	invoiceRepo  repository.Invoice
	userRepo     repository.User
	customerRepo repository.Customer
}

type CreateInvoiceInput struct {
	CustomerID int
	Payment    string
	FeeRate    string
	Deadline   string
}

type InvoiceUsecase interface {
	Create(mail string, in *CreateInvoiceInput) (*model.Invoice, error)
}

func NewInvoiceUsecase(invoiceRepo repository.Invoice, userRepo repository.User, customerRepo repository.Customer) InvoiceUsecase {
	return &invoiceUsecase{
		invoiceRepo:  invoiceRepo,
		userRepo:     userRepo,
		customerRepo: customerRepo,
	}
}

func (u invoiceUsecase) Create(mail string, in *CreateInvoiceInput) (*model.Invoice, error) {
	// ユーザーの属する会社を取得
	user, err := u.userRepo.FindByMail(mail)
	if err != nil {
		return nil, err
	}
	// customerの存在チェック
	if _, err := u.customerRepo.FindByID(in.CustomerID); err != nil {
		return nil, errcode.NewHTTPError(http.StatusBadRequest, "customer is not exist")
	}

	i, err := model.NewInvoice(user.CompanyID, in.CustomerID, in.Payment, in.FeeRate, in.Deadline)
	if err != nil {
		return nil, err
	}
	invoice, err := u.invoiceRepo.Create(i)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}
