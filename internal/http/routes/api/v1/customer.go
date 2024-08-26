package v1

import (
	"crud-customer/config"
	"crud-customer/internal/handler"
	"crud-customer/internal/repository"
	"crud-customer/internal/service"
	"crud-customer/pkg/database"
	"github.com/labstack/echo/v4"
)

func SetCustomerRoutes(cfg *config.Config, echoApp *echo.Echo, db database.GormDB) {
	customerRepo := repository.NewCustomer(db.GetDB(), cfg)
	customerService := service.NewCustomer(cfg, customerRepo)
	customerHandler := handler.NewCustomer(cfg, customerService)
	v1Group := echoApp.Group("/api/v1")
	v1Group.POST("/customers/", customerHandler.CreateCustomer).Name = "CreateCustomer"
	v1Group.PUT("/customers/:id", customerHandler.UpdateCustomer).Name = "UpdateCustomer"
	v1Group.GET("/customers/:id", customerHandler.GetCustomerByID).Name = "GetCustomerByID"
	v1Group.DELETE("/customers/:id", customerHandler.DeleteCustomer).Name = "DeleteCustomer"
	v1Group.GET("/customers/", customerHandler.GetAllCustomer).Name = "GetAllCustomer"
}
