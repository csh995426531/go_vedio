package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}

	defer dbConn.close()

	err = dbConn.Ping()
	if err != nil {
		panic(err.Error())
	}
}
