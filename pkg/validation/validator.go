package validation

import (
	"fmt"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type WithDefaultsValidator struct {
	*validator.Validate
}

// NewWithDefaultsValidator creates a new validator for model fields.
func NewWithDefaultsValidator() WithDefaultsValidator {
	// Create a new validator for a Book model.
	validate := validator.New()
	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field, ok := fl.Field().Interface().(uuid.UUID)

		if !ok {
			if _, err := uuid.Parse(fl.Field().String()); err != nil {
				return false
			}
		}
		value := field.String()

		if _, err := uuid.Parse(value); err != nil {
			return false
		}

		return true
	})

	return WithDefaultsValidator{
		validate,
	}
}

func (v *WithDefaultsValidator) ValidateWithDefaults(dto interface{}) error {
	// set default values
	if err := defaults.Set(dto); err != nil {
		return fmt.Errorf("failed to set default values: %w", err)
	}

	if err := v.Struct(dto); err != nil {
		return fmt.Errorf("failed to validate: %w", err)
	}
	return nil
}
