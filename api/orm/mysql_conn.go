package orm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	conn *sql.DB
	err  error
)

func init() {
	//#@tcp(127.0.0.1:3306)
	conn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/soup?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
