package validators_test

import (
	"testing"
	"transaction/domain/validators"

	"github.com/stretchr/testify/assert"
)

func TestValidator_Validate(t *testing.T) {
	validator := validators.NewValidator()
	validator.Add("field1", validators.ValidateLength(10))
	validator.Add("field2", validators.ValidatePositiveAmount())

	data := map[string]interface{}{
		"field1": "short",
		"field2": 10.0,
	}

	errors := validator.Validate(data)

	assert.Empty(t, errors)
}

func TestValidator_Validate_Error(t *testing.T) {
	validator := validators.NewValidator()
	validator.Add("field1", validators.ValidateLength(10))
	validator.Add("field2", validators.ValidatePositiveAmount())

	data := map[string]interface{}{
		"field1": "this is too long",
		"field2": -10.0,
	}

	errors := validator.Validate(data)

	assert.NotEmpty(t, errors)
	assert.Contains(t, errors[0].Error(), "field1")
	assert.Contains(t, errors[1].Error(), "field2")
}

func TestValidateLength(t *testing.T) {
	rule := validators.ValidateLength(10)

	err := rule("short")
	assert.NoError(t, err)

	err = rule("this is too long")
	assert.Error(t, err)
}

func TestValidatePositiveAmount(t *testing.T) {
	rule := validators.ValidatePositiveAmount()

	err := rule(10.0)
	assert.NoError(t, err)

	err = rule(-10.0)
	assert.Error(t, err)
}
