package query

import (
	"fmt"
	"strings"
	"sync"
	"time"

	model "ai-project/models"
)

func GenerateSQL(query *model.QueryParams, selectStatement, table string) *string {

	var sqlWhereString string

	if query.Filter {

		sqlWhereString = "WHERE"

		for i, group := range *query.Groups {

			var temSQLString string

			for j, filter := range *group.Filters {

				temSQLString = fmt.Sprintf(
					"%v %v %v %v ",
					temSQLString,
					value(convertField(filter.Field, filter.DataType)),
					covertConditionsToSQL(filter.FilterOption),
					convertValue(filter),
				)

				if (j + 1) != len(*group.Filters) {
					temSQLString = fmt.Sprintf(
						"%v %v",
						temSQLString,
						group.FilterSearchCondition,
					)
				}
			}

			if (i + 1) != len(*query.Groups) {
				temSQLString = fmt.Sprintf(
					"(%v) %v",
					temSQLString,
					group.FilterGroupCondition,
				)
			} else {
				temSQLString = fmt.Sprintf("(%v)", temSQLString)
			}

			sqlWhereString = fmt.Sprintf("%v %v", sqlWhereString, temSQLString)
		}

	}

	selectStatement = strings.ReplaceAll(selectStatement, "--WHERE", sqlWhereString)

	selectStatement = strings.ReplaceAll(selectStatement, "--ORDER_BY", fmt.Sprintf("ORDER BY %v %v", value(convertField(query.OrderBy, "")), strings.ToUpper(query.Sort)))

	selectStatement = strings.ReplaceAll(selectStatement, "--OFFSET", fmt.Sprintf("OFFSET (%v - 1) * %v", query.PageNumber, query.BatchNumber))

	selectStatement = strings.ReplaceAll(selectStatement, "--LIMIT", fmt.Sprintf("LIMIT %v", query.BatchNumber))

	result := selectStatement + ";"

	return &result
}

func FilterValidation(source string, data *model.QueryParams) error {
	if data.Filter {

		if len(*data.Groups) == 0 {
			return fmt.Errorf("filter is set to true, groups cannot be empty")
		}

		var wg sync.WaitGroup
		ch := make(chan error)

		for _, v := range *data.Groups {
			wg.Add(1)

			go func(v model.FilterGroup) {
				defer wg.Done()

				if v.Filters != nil {
					for _, f := range *v.Filters {
						wg.Add(1)

						go func(f model.Filter, ch *chan error) {
							defer wg.Done()

							if f.DataType != "string" && f.DataType != "number" && f.DataType != "bool" && f.DataType != "date" && f.DataType != "dateTime" {
								*ch <- fmt.Errorf("dataType can either be string | number | bool | date | dateTime")
							}

							if f.DataType == "date" && !dateQueryValidator(&f.Value) {
								*ch <- fmt.Errorf("date should be in this format 2006-01-02 YYYY-MM-DD YYYY(Year) MM(Month) DD(Day).")
							}

							if f.DataType == "dateTime" && !dateTimeQueryValidator(&f.Value) {
								*ch <- fmt.Errorf("dateTime should be in this format 2006-01-02T15:04:05-07:00 YYYY(Year)='2006', MM(Month)='01', DD(Day)='02', HH(Hour 1 - 24)='15', MM(Minutes)='04', and UTC/GTM TIMEZONE(plus or minus HH:MM) = '+00:00'. a date timestamp with timezone")
							}

							if source == "user" {
								if err := QueryUserFieldValidation(f.Field); err != nil {
									*ch <- err
								}
							}

							if source == "prediction" {
								if err := QueryPredictionFieldValidation(f.Field); err != nil {
									*ch <- err
								}
							}
						}(f, &ch)
					}
				}
			}(v)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()

		for v := range ch {
			return v
		}
	}
	return nil
}

func dateQueryValidator(dateTime *string) bool {
	_, err := time.Parse(time.DateOnly, *dateTime)
	return err == nil
}

func dateTimeQueryValidator(dateTimeZone *string) bool {

	t, err := time.Parse("2006-01-02T15:04:05-07:00", *dateTimeZone)
	if err != nil {
		return err == nil
	}

	zone := t.Format("-07:00")

	if strings.Contains(zone, "-12:") {
		return true
	} else if strings.Contains(zone, "-11:") {
		return true
	} else if strings.Contains(zone, "-10:") {
		return true
	} else if strings.Contains(zone, "-09:") {
		return true
	} else if strings.Contains(zone, "-08:") {
		return true
	} else if strings.Contains(zone, "-07:") {
		return true
	} else if strings.Contains(zone, "-06:") {
		return true
	} else if strings.Contains(zone, "-05:") {
		return true
	} else if strings.Contains(zone, "-04:") {
		return true
	} else if strings.Contains(zone, "-03:") {
		return true
	} else if strings.Contains(zone, "-02:") {
		return true
	} else if strings.Contains(zone, "-01:") {
		return true
	} else if strings.Contains(zone, "+00:") {
		return true
	} else if strings.Contains(zone, "+01:") {
		return true
	} else if strings.Contains(zone, "+02:") {
		return true
	} else if strings.Contains(zone, "+03:") {
		return true
	} else if strings.Contains(zone, "+04:") {
		return true
	} else if strings.Contains(zone, "+05:") {
		return true
	} else if strings.Contains(zone, "+06:") {
		return true
	} else if strings.Contains(zone, "+07:") {
		return true
	} else if strings.Contains(zone, "+08:") {
		return true
	} else if strings.Contains(zone, "+09:") {
		return true
	} else if strings.Contains(zone, "+10:") {
		return true
	} else if strings.Contains(zone, "+11:") {
		return true
	} else if strings.Contains(zone, "+12:") {
		return true
	} else if strings.Contains(zone, "+13:") {
		return true
	} else if strings.Contains(zone, "+14") {
		return true
	} else {
		return false
	}
}

func covertConditionsToSQL(condition string) string {
	switch condition {
	case "ncn":
		return "NOT ILIKE"
	case "nsw":
		return "NOT ILIKE"
	case "new":
		return "NOT ILIKE"
	case "cn":
		return "ILIKE"
	case "sw":
		return "ILIKE"
	case "ew":
		return "ILIKE"
	case "gte":
		return ">="
	case "gt":
		return ">"
	case "lte":
		return "<="
	case "lt":
		return "<"
	case "neq":
		return "<>"
	default:
		return "="
	}
}

func convertField(field, t string) string {
	if strings.ToLower(t) == "date" {
		field = fmt.Sprintf("%v::date", field)
	}

	switch {
	case strings.Contains(field, "prediction."):
		return strings.Replace(field, "prediction.", "P.", 1)
	case strings.Contains(field, "user."):
		return strings.Replace(field, "user.", "U.", 1)
	default:
		return strings.Replace(field, field, field, 1)
	}
}

func value(s string) string {
	switch {
	case strings.Contains(s, "firstName"):
		return strings.Replace(s, "firstName", "first_name", 1)
	case strings.Contains(s, "lastName"):
		return strings.Replace(s, "lastName", "last_name", 1)
	case strings.Contains(s, "groupId"):
		return strings.Replace(s, "groupId", "group_id", 1)
	case strings.Contains(s, "predictionModel"):
		return strings.Replace(s, "predictionModel", "prediction_model", 1)
	case strings.Contains(s, "isVerified"):
		return strings.Replace(s, "isVerified", "is_verified", 1)
	case strings.Contains(s, "createdAt"):
		return strings.Replace(s, "createdAt", "created_at", 1)
	case strings.Contains(s, "updatedAt"):
		return strings.Replace(s, "updatedAt", "updated_at", 1)
	case strings.Contains(s, "expireAt"):
		return strings.Replace(s, "expireAt", "expire_at", 1)
	default:
		return strings.Replace(s, s, s, 1)
	}
}

func convertValue(f model.Filter) string {

	if strings.ToLower(f.FilterOption) == "cn" || strings.ToLower(f.FilterOption) == "ncn" {
		f.Value = fmt.Sprintf("%v%v%v", "%", f.Value, "%")
	}

	if strings.ToLower(f.FilterOption) == "sw" || strings.ToLower(f.FilterOption) == "nsw" {
		f.Value = fmt.Sprintf("%v%v", f.Value, "%")
	}

	if strings.ToLower(f.FilterOption) == "ew" || strings.ToLower(f.FilterOption) == "new" {
		f.Value = fmt.Sprintf("%v%v", "%", f.Value)
	}

	if strings.ToLower(f.DataType) == "string" || strings.ToLower(f.DataType) == "date" {
		f.Value = fmt.Sprintf("'%v'", f.Value)
	}

	if f.DataType == "dateTime" {
		f.Value = fmt.Sprintf("'%v'::TIMESTAMPTZ", f.Value)
	}

	return f.Value
}

func stringReplaceAfter(input, pattern string) string {
	index := strings.Index(input, pattern)
	if index == -1 {
		return input
	}
	return input[:index]
}
