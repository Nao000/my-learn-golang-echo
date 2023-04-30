package core

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"github.com/go-sql-driver/mysql"
)

func DatabseConnect() *sql.DB {

	var db *sql.DB

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "db:3306",
		DBName: "employees",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Database Connected!")

	return db
}
