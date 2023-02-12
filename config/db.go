package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	fmt.Println("Connecting to Database")

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_USER_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
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
