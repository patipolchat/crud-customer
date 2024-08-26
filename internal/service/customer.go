package service

import (
	"context"
	"crud-customer/internal/entity"
)

type Customer interface {
	CreateCustomer(ctx context.Context, name string, age uint) (*entity.Customer, error)
	UpdateCustomer(ctx context.Context, id uint, customer *entity.Customer) (*entity.Customer, error)
	GetCustomerByID(ctx context.Context, id uint) (*entity.Customer, error)
	DeleteCustomer(ctx context.Context, id uint) error
	GetAllCustomer(ctx context.Context) ([]*entity.Customer, error)
}
