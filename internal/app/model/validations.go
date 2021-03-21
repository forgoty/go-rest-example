package model

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

func requiredIf(condition bool) validation.RuleFunc {
	return func(value interface{}) error {
		if condition {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
