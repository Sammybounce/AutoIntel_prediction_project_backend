package model

import (
	"database/sql"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

type UserFull struct {
	*User

	UpdatedAt string `json:"updatedAt"`
	Deleted   bool   `json:"deleted"`
	Password  string `json:"password"`
}

type UserSQL struct {
	Id        sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
	Email     sql.NullString
	CreatedAt sql.NullTime
}

type UserFullSQL struct {
	*UserSQL

	UpdatedAt sql.NullTime
	Deleted   sql.NullBool
	Password  sql.NullString
}
