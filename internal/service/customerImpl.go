package service

import (
	"context"
	"crud-customer/config"
	"crud-customer/internal/entity"
	"crud-customer/internal/repository"
)

type customerImpl struct {
	customerRepo repository.Customer
	cfg          *config.Config
}

func (c *customerImpl) CreateCustomer(ctx context.Context, name string, age uint) (*entity.Customer, error) {
	customer := &entity.Customer{
		Name: &name,
		Age:  &age,
	}

	id, err := c.customerRepo.CreateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}
	customer.ID = *id
	return customer, nil
}

func (c *customerImpl) UpdateCustomer(ctx context.Context, id uint, customer *entity.Customer) (*entity.Customer, error) {
	customer, err := c.customerRepo.UpdateCustomer(ctx, id, customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *customerImpl) GetCustomerByID(ctx context.Context, id uint) (*entity.Customer, error) {
	customer, err := c.customerRepo.GetCustomerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *customerImpl) DeleteCustomer(ctx context.Context, id uint) error {
	return c.customerRepo.DeleteCustomer(ctx, id)
}

func (c *customerImpl) GetAllCustomer(ctx context.Context) ([]*entity.Customer, error) {
	customers, err := c.customerRepo.GetAllCustomer(ctx)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func NewCustomer(cfg *config.Config, customerRepo repository.Customer) Customer {
	return &customerImpl{
		customerRepo: customerRepo,
		cfg:          cfg,
	}
}
