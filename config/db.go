package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	HOST     = "127.0.0.1"
	PORT     = 3306
	USER     = "root"
	PASSWORD = ""
	DBNAME   = "cake"
)

func NewDB() *sql.DB {
	fmt.Println("Connecting to Database")

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		USER, PASSWORD, HOST, PORT, DBNAME,
	)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Println("Unable to connect to database ", err.Error())
		return nil
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Unable to connect to database ", err.Error())
		return nil
	}

	fmt.Println("Database connected!")

	return db

}
