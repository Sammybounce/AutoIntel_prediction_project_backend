package query

import (
	"errors"
	"fmt"

	"ai-project/utils/array"
)

type AllowedSearchConditionStruct struct {
	Condition string `json:"condition"`
	Meaning   string `json:"meaning"`
}

var AllowedSearchCondition = []AllowedSearchConditionStruct{
	{
		Condition: "ncn",
		Meaning:   "does not contain",
	},
	{
		Condition: "new",
		Meaning:   "does not end with",
	},
	{
		Condition: "nsw",
		Meaning:   "does not start with",
	},
	{
		Condition: "cn",
		Meaning:   "contain",
	},

	{
		Condition: "sw",
		Meaning:   "start with",
	},

	{
		Condition: "ew",
		Meaning:   "end with",
	},

	{
		Condition: "gt",
		Meaning:   "greater then",
	},

	{
		Condition: "gte",
		Meaning:   "greater then or equals to",
	},

	{
		Condition: "lt",
		Meaning:   "less then",
	},

	{
		Condition: "lte",
		Meaning:   "less then or equals to",
	},

	{
		Condition: "eq",
		Meaning:   "equals to",
	},
	{
		Condition: "neq",
		Meaning:   "not equals to",
	},
}

func QuerySearchConditionValidation(field string) error {

	_, check, _ := array.Find(&AllowedSearchCondition, func(d AllowedSearchConditionStruct) bool {
		return field == d.Condition
	})

	if !check {
		err := fmt.Sprintf("%v is not allowed go to /allowed-query/search-condition to see the list of allowed query fields", field)
		return errors.New(err)
	}

	return nil
}
