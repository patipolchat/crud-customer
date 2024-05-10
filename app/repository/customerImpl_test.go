package repository

import (
	"context"
	"crud-customer/config"
	"crud-customer/entity"
	"crud-customer/util/typehelper"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"testing"
)

type CustomerImplTestSuite struct {
	suite.Suite
	customer  Customer
	tmpDBFile *os.File
	db        *gorm.DB
	tx        *gorm.DB
}

func (s *CustomerImplTestSuite) SetupSuite() {
	f, err := os.CreateTemp("", "test.*.db")
	if err != nil {
		panic(err)
	}
	s.tmpDBFile = f
	db, err := gorm.Open(sqlite.Open(f.Name()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	s.db = db
	if err := s.db.AutoMigrate(&entity.Customer{}); err != nil {
		panic(err)
	}
}

func (s *CustomerImplTestSuite) TearDownSuite() {
	os.Remove(s.tmpDBFile.Name())
	s.db = nil
}

func (s *CustomerImplTestSuite) SetupTest() {
	s.tx = s.db.Begin()
	s.customer = NewCustomer(s.tx, &config.Config{})
}

func (s *CustomerImplTestSuite) TearDownTest() {
	s.tx.Rollback()
	s.customer = nil
}

func (s *CustomerImplTestSuite) TestCreateCustomerSuccess() {
	args := &entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	}

	want := typehelper.GetPointer(uint(1))

	got, err := s.customer.CreateCustomer(context.Background(), args)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *CustomerImplTestSuite) TestCreateCustomerError() {
	args := &entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  nil,
	}

	got, err := s.customer.CreateCustomer(context.Background(), args)
	s.Error(err)
	s.Nil(got)
}

func (s *CustomerImplTestSuite) TestUpdateCustomerSuccess() {
	result := s.tx.Create(&entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	})
	if result.Error != nil {
		panic(result.Error)
	}

	want := &entity.Customer{
		ID:   1,
		Name: typehelper.GetPointer("John Dee"),
		Age:  typehelper.GetPointer(uint(20)),
	}

	got, err := s.customer.UpdateCustomer(context.Background(), 1, &entity.Customer{
		Name: typehelper.GetPointer("John Dee"),
		Age:  typehelper.GetPointer(uint(20)),
	})
	s.NoError(err)
	s.Equal(want, got)
}

func (s *CustomerImplTestSuite) TestUpdateCustomerError() {
	result := s.tx.Create(&entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	})
	if result.Error != nil {
		panic(result.Error)
	}

	got, err := s.customer.UpdateCustomer(context.Background(), 1, &entity.Customer{
		Name: typehelper.GetPointer("John Dee"),
		Age:  nil,
	})
	s.Error(err)
	s.Nil(got)
}

func (s *CustomerImplTestSuite) TestGetCustomerByIDSuccess() {
	result := s.tx.Create(&entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	})
	if result.Error != nil {
		panic(result.Error)
	}

	want := &entity.Customer{
		ID:   1,
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	}

	got, err := s.customer.GetCustomerByID(context.Background(), 1)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *CustomerImplTestSuite) TestGetCustomerByIDError() {
	got, err := s.customer.GetCustomerByID(context.Background(), 1)
	s.Error(err)
	s.Nil(got)
}

func (s *CustomerImplTestSuite) TestDeleteCustomerSuccess() {
	result := s.tx.Create(&entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	})
	if result.Error != nil {
		panic(result.Error)
	}

	err := s.customer.DeleteCustomer(context.Background(), 1)
	s.NoError(err)
}

func (s *CustomerImplTestSuite) TestGetAllCustomerSuccess() {
	result := s.tx.Create(&entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	})
	if result.Error != nil {
		panic(result.Error)
	}

	want := []*entity.Customer{
		{
			ID:   1,
			Name: typehelper.GetPointer("John Doe"),
			Age:  typehelper.GetPointer(uint(20)),
		},
	}

	got, err := s.customer.GetAllCustomer(context.Background())
	s.NoError(err)
	s.Equal(want, got)
}

func TestCustomerImplSuite(t *testing.T) {
	suite.Run(t, new(CustomerImplTestSuite))
}
