package repository

import (
	"database/sql"
)

var Db *sql.DB

func init() {
	var err error
	connStr := "host=sample-api-db user=todo password=root sslmode=disable"
	Db, err = sql.Open("postgres", connStr)
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	// 	"todo-app", "todo-password", "sample-api-db:3306", "todo",
	// )
	// Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
}