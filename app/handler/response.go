package handler

import (
	"github.com/labstack/echo/v4"
)

type CustomerData struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type CreateUpdateCustomerResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    CustomerData `json:"data"`
}

type DeleteCustomerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetAllCustomerResponse struct {
	Success bool           `json:"success"`
	Data    []CustomerData `json:"data"`
	Message string         `json:"message"`
}

type GetCustomerByIDResponse struct {
	Success bool         `json:"success"`
	Data    CustomerData `json:"data"`
	Message string       `json:"message"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

func NewErrorResponse(statusCode int, message string) error {
	resp := &ErrorResponse{
		StatusCode: statusCode,
		Success:    false,
		Message:    message,
	}
	return echo.NewHTTPError(statusCode, resp)
}
