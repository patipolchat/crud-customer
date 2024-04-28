package repository

import (
	"context"
	"crud-customer/config"
	"crud-customer/entity"
	"gorm.io/gorm"
)

type customerImpl struct {
	db  *gorm.DB
	cfg *config.Config
}

func (c *customerImpl) CreateCustomer(ctx context.Context, customer *entity.Customer) (*uint, error) {
	result := c.db.WithContext(ctx).Create(customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer.ID, nil
}

func (c *customerImpl) UpdateCustomer(ctx context.Context, id uint, customer *entity.Customer) (*entity.Customer, error) {
	result := c.db.WithContext(ctx).Raw("UPDATE customers set name = ?, age = ? where id = ? RETURNING *", customer.Name, customer.Age, id).Scan(&customer)

	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

func (c *customerImpl) GetCustomerByID(ctx context.Context, id uint) (*entity.Customer, error) {
	var customer entity.Customer
	result := c.db.WithContext(ctx).First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (c *customerImpl) DeleteCustomer(ctx context.Context, id uint) error {
	result := c.db.WithContext(ctx).Delete(&entity.Customer{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *customerImpl) GetAllCustomer(ctx context.Context) ([]*entity.Customer, error) {
	var customers []*entity.Customer
	result := c.db.WithContext(ctx).Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func NewCustomer(db *gorm.DB, cfg *config.Config) Customer {
	return &customerImpl{
		db:  db,
		cfg: cfg,
	}
}
