package validator

import (
	govalidator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

var (
	echoValidatorOnce     sync.Once
	echoValidatorInstance *EchoValidator
)

type EchoValidator struct {
	validator *govalidator.Validate
}

func (e *EchoValidator) Validate(i interface{}) error {
	if err := e.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func GetEchoValidator() *EchoValidator {
	echoValidatorOnce.Do(func() {
		echoValidatorInstance = &EchoValidator{GetValidator()}
	})
	return echoValidatorInstance
}
