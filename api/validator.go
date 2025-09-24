package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/w00ye0l/simplebank/util"
)

var validCurrency validator.Func = func(fildLevel validator.FieldLevel) bool {
	if currency, ok := fildLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
