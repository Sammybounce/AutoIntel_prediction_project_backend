package model

import (
	"database/sql"
	"time"
)

type UserToken struct {
	Id       string `json:"id"`
	UserId   string `json:"userId"`
	Token    string `json:"token"`
	ExpireAt string `json:"expireAt"`
}

type UserTokenFull struct {
	*UserToken

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Deleted   bool      `json:"deleted"`
}

type UserTokenSQL struct {
	Id       sql.NullString
	UserId   sql.NullString
	Token    sql.NullString
	ExpireAt sql.NullTime
}

type UserTokenFullSQL struct {
	*UserTokenFullSQL

	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Deleted   sql.NullBool
}
