package structValidator

import (
	"fmt"
	"reflect"
	"strings"

	model "ai-project/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Validate(data interface{}) error {
	validate = validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(data)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var message string

		for _, err := range err.(validator.ValidationErrors) {
			e := model.ValidationError{
				Namespace:       err.Namespace(),
				Field:           err.Field(),
				StructNamespace: err.StructNamespace(),
				StructField:     err.StructField(),
				Tag:             err.Tag(),
				ActualTag:       err.ActualTag(),
				Kind:            fmt.Sprintf("%v", err.Kind()),
				Type:            fmt.Sprintf("%v", err.Type()),
				Value:           fmt.Sprintf("%v", err.Value()),
				Param:           err.Param(),
				Message:         err.Error(),
			}

			message = fmt.Sprintf("%v %v", message, e.Message)
		}

		// from here you can create your own error messages in whatever language you wish
		return fmt.Errorf("%v", message)
	}
	return nil
}
