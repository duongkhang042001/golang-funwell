package validator

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(ctx context.Context, s interface{}) error {
	err := validate.StructCtx(ctx, s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return mappingValidationErrorMessages(validationErrors)
		}
		return err
	}
	return nil
}

func mappingValidationErrorMessages(errs validator.ValidationErrors) error {
	var messages []string
	for _, e := range errs {
		field := e.Field()
		tag := e.Tag()
		param := e.Param()

		switch tag {
		case "required":
			messages = append(messages, fmt.Sprintf("Field '%s' is required.", field))
		case "len":
			messages = append(messages, fmt.Sprintf("Field '%s' must have exactly %s characters.", field, param))
		case "min":
			messages = append(messages, fmt.Sprintf("Field '%s' must have at least %s characters.", field, param))
		case "max":
			messages = append(messages, fmt.Sprintf("Field '%s' must have no more than %s characters.", field, param))
		case "email":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid email address.", field))
		case "uuid":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid UUID.", field))
		case "gte":
			messages = append(messages, fmt.Sprintf("Field '%s' must be greater than or equal to %s.", field, param))
		case "lte":
			messages = append(messages, fmt.Sprintf("Field '%s' must be less than or equal to %s.", field, param))
		case "alpha":
			messages = append(messages, fmt.Sprintf("Field '%s' must contain only alphabetic characters.", field))
		case "numeric":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a numeric value.", field))
		case "alphanum":
			messages = append(messages, fmt.Sprintf("Field '%s' must contain only alphabetic and numeric characters.", field))
		case "alphaunicode":
			messages = append(messages, fmt.Sprintf("Field '%s' must contain only Unicode alphabetic characters.", field))
		case "gt":
			messages = append(messages, fmt.Sprintf("Field '%s' must be greater than %s.", field, param))
		case "lt":
			messages = append(messages, fmt.Sprintf("Field '%s' must be less than %s.", field, param))
		case "datetime":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid datetime with format %s.", field, param))
		case "json":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid JSON string.", field))
		case "hostname":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid hostname.", field))
		case "isbn":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid ISBN.", field))
		case "isbn10":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid ISBN-10.", field))
		case "isbn13":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid ISBN-13.", field))
		case "hexcolor":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid hex color code.", field))
		case "color":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid color.", field))
		case "cpf":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid CPF number.", field))
		case "cnpj":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid CNPJ number.", field))
		case "passport":
			messages = append(messages, fmt.Sprintf("Field '%s' must be a valid passport number.", field))
		default:
			messages = append(messages, fmt.Sprintf("Field '%s' has an invalid value.", field))
		}
	}

	return fmt.Errorf(strings.Join(messages, " "))
}
