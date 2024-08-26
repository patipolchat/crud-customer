package service

import (
	"context"
	"crud-customer/config"
	"crud-customer/internal/entity"
	mockrepo "crud-customer/mocks/internal_/repository"
	"crud-customer/util/typehelper"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CustomerImplTestSuite struct {
	suite.Suite
	mockCustomerRepo *mockrepo.Customer
	customer         Customer
}

func (s *CustomerImplTestSuite) TearDownTest() {
	s.mockCustomerRepo = nil
	s.customer = nil
}

func (s *CustomerImplTestSuite) SetupTest() {
	s.mockCustomerRepo = mockrepo.NewCustomer(s.T())
	s.customer = NewCustomer(&config.Config{}, s.mockCustomerRepo)
}

func (s *CustomerImplTestSuite) TestCreateCustomerSuccess() {
	s.mockCustomerRepo.EXPECT().CreateCustomer(mock.Anything, &entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	}).Return(typehelper.GetPointer(uint(1)), nil)

	want := &entity.Customer{
		ID:   1,
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	}

	got, err := s.customer.CreateCustomer(context.Background(), "John Doe", 20)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *CustomerImplTestSuite) TestCreateCustomerError() {
	s.mockCustomerRepo.EXPECT().CreateCustomer(mock.Anything, &entity.Customer{
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	}).Return(nil, fmt.Errorf("error"))

	got, err := s.customer.CreateCustomer(context.Background(), "John Doe", 20)
	s.Error(err)
	s.Nil(got)
}

func (s *CustomerImplTestSuite) TestUpdateCustomerSuccess() {
	customer := &entity.Customer{
		ID:   1,
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	}

	s.mockCustomerRepo.EXPECT().UpdateCustomer(mock.Anything, uint(1), customer).Return(customer, nil)

	want := customer

	got, err := s.customer.UpdateCustomer(context.Background(), 1, customer)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *CustomerImplTestSuite) TestUpdateCustomerError() {
	customer := &entity.Customer{
		ID:   1,
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	}

	s.mockCustomerRepo.EXPECT().UpdateCustomer(mock.Anything, uint(1), customer).Return(nil, fmt.Errorf("error"))

	got, err := s.customer.UpdateCustomer(context.Background(), 1, customer)
	s.Error(err)
	s.Nil(got)
}

func (s *CustomerImplTestSuite) TestGetCustomerByIDSuccess() {
	want := &entity.Customer{
		ID:   1,
		Name: typehelper.GetPointer("John Doe"),
		Age:  typehelper.GetPointer(uint(20)),
	}

	s.mockCustomerRepo.EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(want, nil)

	got, err := s.customer.GetCustomerByID(context.Background(), 1)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *CustomerImplTestSuite) TestGetCustomerByIDError() {
	s.mockCustomerRepo.EXPECT().GetCustomerByID(mock.Anything, uint(1)).Return(nil, fmt.Errorf("error"))

	got, err := s.customer.GetCustomerByID(context.Background(), 1)
	s.Error(err)
	s.Nil(got)
}

func (s *CustomerImplTestSuite) TestDeleteCustomerSuccess() {
	s.mockCustomerRepo.EXPECT().DeleteCustomer(mock.Anything, uint(1)).Return(nil)

	err := s.customer.DeleteCustomer(context.Background(), 1)
	s.NoError(err)
}

func (s *CustomerImplTestSuite) TestDeleteCustomerError() {
	s.mockCustomerRepo.EXPECT().DeleteCustomer(mock.Anything, uint(1)).Return(fmt.Errorf("error"))

	err := s.customer.DeleteCustomer(context.Background(), 1)
	s.Error(err)
}

func (s *CustomerImplTestSuite) TestGetAllCustomerSuccess() {
	want := []*entity.Customer{
		{
			ID:   1,
			Name: typehelper.GetPointer("John Doe"),
			Age:  typehelper.GetPointer(uint(20)),
		},
	}

	s.mockCustomerRepo.EXPECT().GetAllCustomer(mock.Anything).Return(want, nil)

	got, err := s.customer.GetAllCustomer(context.Background())
	s.NoError(err)
	s.Equal(want, got)
}

func (s *CustomerImplTestSuite) TestGetAllCustomerError() {
	s.mockCustomerRepo.EXPECT().GetAllCustomer(mock.Anything).Return(nil, fmt.Errorf("error"))

	got, err := s.customer.GetAllCustomer(context.Background())
	s.Error(err)
	s.Nil(got)
}

func TestCustomerImplTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerImplTestSuite))
}
