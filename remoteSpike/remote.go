package remoteSpike
//
//import (
//	sql "database/sql"
//	_ "github.com/go-sql-driver/mysql"
//)
//
//var DB *sql.DB
//
//func NewClient()(db *sql.DB,err error){
//	db,err = sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/ms")
//	if err != nil {
//		return
//	}
//	DB = db
//	return db,err
//}