package model

type ResponseWithSingleData struct {
	Data    interface{} `json:"data"`
	Type    string      `json:"type"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

type ResponseWithMultipleData struct {
	Data         interface{} `json:"data"`
	Type         string      `json:"type"`
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	PageNumber   int         `json:"pageNumber"`
	BatchNumber  int         `json:"batchNumber"`
	TotalRecords int         `json:"totalRecords"`
}

type QueryParams struct {
	More        bool           `json:"more"`
	Filter      bool           `json:"filter"`
	PageNumber  int            `json:"pageNumber" validate:"required"`
	BatchNumber int            `json:"batchNumber" validate:"required"`
	Sort        string         `json:"sort" validate:"required"`
	OrderBy     string         `json:"orderBy" validate:"required"`
	Groups      *[]FilterGroup `json:"groups" validate:"dive"`
}

type FilterGroup struct {
	FilterGroupCondition  string    `json:"filterGroupCondition" validate:"required"`
	FilterSearchCondition string    `json:"filterSearchCondition" validate:"required"`
	Filters               *[]Filter `json:"filters" validate:"gt=0,dive,required"`
}

type Filter struct {
	FilterOption string `json:"filterOption" validate:"required"`
	Field        string `json:"field" validate:"required"`
	DataType     string `json:"dataType" validate:"required"`
	Value        string `json:"value"`
}
