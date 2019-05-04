package orm

import "database/sql"

var dbConn *sql.DB
var err error

func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/soup?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
