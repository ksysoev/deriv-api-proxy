package validator

import "fmt"

type ValidatorConfig struct {
	fields map[string]*FielConfig
}

type FielConfig struct {
	Type string `mapstructure:"type"`
}

type validator struct {
	config *ValidatorConfig
}

func New(config *ValidatorConfig) (*validator, error) {
	for field, fieldConfig := range config.fields {
		if fieldConfig.Type != "string" && fieldConfig.Type != "number" && fieldConfig.Type != "bool" {
			return nil, fmt.Errorf("unknown type %s for field %s", fieldConfig.Type, field)
		}
	}

	return &validator{config: config}, nil
}

func (v *validator) Validate(data map[string]any) error {
	errValidation := NewValidationError()

	for field, config := range v.config.fields {
		if value, ok := data[field]; ok {
			if err := v.validateField(config, value); err != nil {
				errValidation.AddError(field, err)
			}
		} else {
			errValidation.AddError(field, fmt.Errorf("field %s is missing", field))
		}
	}

	for field := range data {
		if _, ok := v.config.fields[field]; !ok {
			errValidation.AddError(field, fmt.Errorf("field %s is not allowed", field))
		}
	}

	if errValidation.HasErrors() {
		return errValidation
	}

	return nil
}

func (v *validator) validateField(config *FielConfig, value any) error {
	switch config.Type {
	case "string":
		if _, ok := value.(string); !ok {
			return fmt.Errorf("field %s is not a string", value)
		}
	case "number":
		if _, ok := value.(float64); !ok {
			return fmt.Errorf("field %s is not an int", value)
		}
	case "bool":
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("field %s is not a bool", value)
		}
	default:
		return fmt.Errorf("unknown field type %s", config.Type)
	}

	return nil
}
