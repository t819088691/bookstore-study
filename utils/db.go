package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:QWERasdfZXCV@tcp(192.168.24.27:30884)/bookstore")

	if err != nil {
		panic(err.Error())
	}

}
