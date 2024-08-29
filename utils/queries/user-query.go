package query

import (
	"errors"
	"fmt"

	"ai-project/utils/array"
)

var AllowedUserFields = []string{
	"user.id",
	"user.firstName",
	"user.lastName",
	"user.email",
	"user.createdAt",
	"user.updatedAt",
	"user.deleted",
}

func QueryUserFieldValidation(field string) error {

	_, check, _ := array.Find[string](&AllowedUserFields, func(d string) bool {
		return d == field
	})

	if !check {
		err := fmt.Sprintf("%v is not allowed go to /allowed-query/fields/users to see the list of allowed query fields", field)
		return errors.New(err)
	}

	return nil
}
