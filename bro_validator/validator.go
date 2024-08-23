package bro_validator

import "github.com/go-playground/validator/v10"

var validatorInstance *validator.Validate

func init() {
	validatorInstance = validator.New(validator.WithRequiredStructEnabled())
}

func Struct(s any) error {
	return validatorInstance.Struct(s)
}
