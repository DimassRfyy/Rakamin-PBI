package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB, err error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "evermos"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
