package data

import (
	"database/sql"
	"time"
)

const dbTimeOut = time.Second * 3

var db *sql.DB

type Models struct {
	User User
}

func New(dbpool *sql.DB) Models {
	db = dbpool
	return Models{
		User: User{},
	}
}
