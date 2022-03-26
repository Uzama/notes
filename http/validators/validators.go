package validators

import (
	"context"
	"errors"
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator  *validator.Validate
	translator ut.Translator
}

func NewValidator() Validator {
	validator := Validator{
		validator: validator.New(),
	}

	return validator
}

func (val Validator) Validate(ctx context.Context, data interface{}) error {

	err := val.validator.Struct(data)
	if err == nil {
		return nil
	}

	fieldErrors := err.(validator.ValidationErrors)

	errMessage := ""

	for i, e := range fieldErrors {
		errMessage += fmt.Sprintf("param:%v is invalid. %s", e.Value(), e.Translate(val.translator))

		if i == len(fieldErrors)-1 {
			continue
		}

		errMessage += " | "
	}

	return errors.New(errMessage)
}
