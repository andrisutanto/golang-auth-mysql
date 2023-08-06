package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//next ganti ke env, agar tidak terekspose di file golangnya
func DBConn() (db *sql.DB, err error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go_db"

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
