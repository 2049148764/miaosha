package model

import (
	sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewClient()(err error){
	DB,err = sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/ms")
	if err != nil {
		return
	}
	return
}