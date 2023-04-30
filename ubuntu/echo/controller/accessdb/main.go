package accessdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

type Employee struct {
    Emp_no     int64  `json:"emp_no"`
    Birth_date string `json:"birth_date"`
    First_name string `json:"first_name"`
    Last_name  string `json:"last_name"`
    Gender     string `json:"gender"`
    Hire_date  string `json:"hire_date"`
}

func Index(c echo.Context) error {

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "db:3306",
		DBName: "employees",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	employees, err := getAllEmployees()
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusCreated, employees)
}

/*
* DBからデータ取得
*/
func getAllEmployees() ([]Employee, error) {

	var employees []Employee

	rows, err := db.Query("SELECT emp_no, first_name, last_name FROM employees WHERE emp_no < 20001", )
	if err != nil {
		return nil, fmt.Errorf("getAllEmployees : %v", err)
	}

	defer rows.Close()

	// １行ずつ取得して構造体を要素に持つスライスに入れる
	for rows.Next() {
		var employee Employee

		// スライスに入れるようの構造体に入れてる
		if err := rows.Scan(&employee.Emp_no, &employee.First_name, &employee.Last_name); err != nil {
			return nil, fmt.Errorf("getAllEmployees : %v", err)
		}

		// ここでスライスに入れてる
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllEmployees %q: %v", err)
	}
	return employees, nil
}