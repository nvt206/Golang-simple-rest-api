package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var m = map[string]string{
	"email":"please provide a valid mail%v",
	"required":"field required%v",
	"min":"min length is %v",
	"max":"max length is %v",
}

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
			Msg: fmt.Sprintf(m[fieldErr.Tag()],fieldErr.Param()),
		}
		rs = append(rs, vErr)
	}
	return rs
}