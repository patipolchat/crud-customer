package handler

import "github.com/labstack/echo/v4"

type Customer interface {
	CreateCustomer(c echo.Context) error
	UpdateCustomer(c echo.Context) error
	GetCustomerByID(c echo.Context) error
	DeleteCustomer(c echo.Context) error
	GetAllCustomer(c echo.Context) error
}
