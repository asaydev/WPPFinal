package util

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func Initializedatabase()*sql.DB  {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/web_project")
	if err != nil {
		log.Print(err.Error())
	}
	return db
}

