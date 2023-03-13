package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateIfContainsMethylParaben(field validator.FieldLevel) bool {
	return !strings.Contains(field.Field().String(), strings.ToLower("methylparaben"))
}
