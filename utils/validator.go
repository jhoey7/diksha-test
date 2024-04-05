package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

// RequiredIfFieldEqual validates field, a field will be mandatory if another field has a specific value
func RequiredIfFieldEqual(fl validator.FieldLevel) bool {
	param := strings.Split(fl.Param(), `:`)
	comparatorField := param[0]
	comparatorValue := param[1]

	var paramFieldValue reflect.Value
	if fl.Parent().Kind() == reflect.Ptr {
		paramFieldValue = fl.Parent().Elem().FieldByName(comparatorField)
	} else {
		paramFieldValue = fl.Parent().FieldByName(comparatorField)
	}

	if paramFieldValue.String() == comparatorValue {
		return fl.Field().String() != ""
	}

	return true
}
