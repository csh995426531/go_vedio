package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //mysql数据库rom驱动
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/vedio?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	// defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		panic(err.Error())
	}
}
