package handler

import (
	"crud-customer/app/service"
	"crud-customer/config"
	"crud-customer/entity"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type customerImpl struct {
	customerService service.Customer
	cfg             *config.Config
}

func (cu *customerImpl) CreateCustomer(c echo.Context) error {
	req := new(CreateCustomerRequest)
	if err := c.Bind(req); err != nil {
		return NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("error binding request: %v", err))
	}
	if err := c.Validate(req); err != nil {
		return NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("error validating request: %v", err))
	}

	customer, err := cu.customerService.CreateCustomer(c.Request().Context(), req.Name, req.Age)
	if err != nil {
		return NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("error creating customer: %v", err))
	}

	resp := &CreateUpdateCustomerResponse{
		Success: true,
		Message: "customer created successfully",
		Data: CustomerData{
			ID:   customer.ID,
			Name: *customer.Name,
			Age:  *customer.Age,
		},
	}

	return c.JSON(http.StatusCreated, resp)
}

func (cu *customerImpl) UpdateCustomer(c echo.Context) error {
	req := new(UpdateCustomerRequest)
	if err := c.Bind(req); err != nil {
		return NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("error binding request: %v", err))
	}
	if err := c.Validate(req); err != nil {
		return NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("error validating request: %v", err))
	}

	_, err := cu.customerService.GetCustomerByID(c.Request().Context(), req.ID)
	if err != nil {
		return NewErrorResponse(http.StatusNotFound, fmt.Sprintf("customer not found: %v", err))
	}

	customer, err := cu.customerService.UpdateCustomer(c.Request().Context(), req.ID, &entity.Customer{
		Name: &req.Name,
		Age:  &req.Age,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error updating customer: %v", err))
	}

	resp := &CreateUpdateCustomerResponse{
		Success: true,
		Message: "customer updated successfully",
		Data: CustomerData{
			ID:   customer.ID,
			Name: *customer.Name,
			Age:  *customer.Age,
		},
	}

	return c.JSON(http.StatusOK, resp)
}

func (cu *customerImpl) DeleteCustomer(c echo.Context) error {
	req := new(DeleteCustomerRequest)
	if err := c.Bind(req); err != nil {
		return NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("error binding request: %v", err))
	}
	if err := c.Validate(req); err != nil {
		return NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("error validating request: %v", err))
	}

	_, err := cu.customerService.GetCustomerByID(c.Request().Context(), req.ID)
	if err != nil {
		return NewErrorResponse(http.StatusNotFound, fmt.Sprintf("customer not found: %v", err))
	}

	err = cu.customerService.DeleteCustomer(c.Request().Context(), req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("error deleting customer: %v", err))
	}

	resp := &DeleteCustomerResponse{
		Success: true,
		Message: "customer deleted successfully",
	}

	return c.JSON(http.StatusOK, resp)
}

func (cu *customerImpl) GetCustomerByID(c echo.Context) error {
	req := new(GetCustomerByIDRequest)
	if err := c.Bind(req); err != nil {
		return NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("error binding request: %v", err))
	}
	if err := c.Validate(req); err != nil {
		return NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("error validating request: %v", err))
	}

	customer, err := cu.customerService.GetCustomerByID(c.Request().Context(), req.ID)
	if err != nil {
		return NewErrorResponse(http.StatusNotFound, fmt.Sprintf("customer not found: %v", err))
	}

	resp := &GetCustomerByIDResponse{
		Success: true,
		Message: "customer found",
		Data: CustomerData{
			ID:   customer.ID,
			Name: *customer.Name,
			Age:  *customer.Age,
		},
	}

	return c.JSON(http.StatusOK, resp)
}

func (cu *customerImpl) GetAllCustomer(c echo.Context) error {
	customers, err := cu.customerService.GetAllCustomer(c.Request().Context())
	if err != nil {
		return NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("error getting all customers: %v", err))
	}

	var data []CustomerData
	for _, customer := range customers {
		data = append(data, CustomerData{
			ID:   customer.ID,
			Name: *customer.Name,
			Age:  *customer.Age,
		})
	}

	resp := &GetAllCustomerResponse{
		Success: true,
		Data:    data,
		Message: "customers found",
	}

	return c.JSON(http.StatusOK, resp)
}

func NewCustomer(cfg *config.Config, customerService service.Customer) Customer {
	return &customerImpl{
		customerService: customerService,
		cfg:             cfg,
	}
}
