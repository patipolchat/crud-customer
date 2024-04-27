package validator

import "sync"

import (
	govalidator "github.com/go-playground/validator/v10"
)

var (
	once              sync.Once
	validatorInstance *govalidator.Validate
)

func GetValidator() *govalidator.Validate {
	once.Do(func() {
		validate := govalidator.New()
		validatorInstance = validate
	})

	return validatorInstance
}
