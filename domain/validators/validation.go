package validators

import (
	"fmt"
)

type Validator struct {
	rules map[string][]Rule
}

func NewValidator() *Validator {
	return &Validator{
		rules: make(map[string][]Rule),
	}
}

func (v *Validator) Add(field string, rule Rule) {
	v.rules[field] = append(v.rules[field], rule)
}

func (v *Validator) Validate(data map[string]interface{}) []error {
	var errors []error
	for field, rules := range v.rules {
		if fieldData, ok := data[field]; ok {
			for _, rule := range rules {
				if err := rule(fieldData); err != nil {
					errors = append(errors, fmt.Errorf("%s: %s", field, err))
				}
			}
		} else {
			errors = append(errors, fmt.Errorf("field %s is missing", field))
		}
	}
	return errors
}

type Rule func(value interface{}) error

func ValidateLength(maxLength int) Rule {
	return func(value interface{}) error {
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("must be a string")
		}
		if len(str) > maxLength {
			return fmt.Errorf("must be less than or equal to %d characters", maxLength)
		}
		return nil
	}
}

func ValidatePositiveAmount() Rule {
	return func(value interface{}) error {
		amount, ok := value.(float64)
		if !ok {
			return fmt.Errorf("must be a number")
		}
		if amount <= 0 {
			return fmt.Errorf("must be greater than 0")
		}
		return nil
	}
}
