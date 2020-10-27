package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)
type ValidateError struct {
	Field string `json:"field"`
	ErrorType string `json:"error_type"`
	Value interface{} `json:"value"`
	Msg string `json:"msg"`
}

func GetErrors(errors validator.ValidationErrors) []ValidateError{
	var rs []ValidateError
	for _,fieldErr := range errors {
		vErr := ValidateError{
			Field: fieldErr.Field(),
			ErrorType: fieldErr.ActualTag(),
			Value: fieldErr.Value(),
			Msg: GetMessage(fieldErr.Tag(),fieldErr.Param()),
		}
		rs = append(rs, vErr)
	}
	return rs
}

func GetMessage(tag string,attrs ...interface{}) string{
	result :=""

	switch tag {
		case "email":
			result = "please provide a valid mail"
		case "required":
			result = "field required"
		case "min":
			result = fmt.Sprintf("min length is %v",attrs)
		case "max":
			result = fmt.Sprintf("max length is %v",attrs)
	}
	return result
}