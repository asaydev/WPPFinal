package util

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func Initializedatabase()*sql.DB  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/web_project")
	if err != nil {
		log.Print(err.Error())
	}
	return db
}

