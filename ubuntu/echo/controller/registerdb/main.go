package registerdb

import (
	"fmt"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"echo/core"
)

// curl -X POST http://localhost/registerdb で試せる
func Index(c echo.Context) error {
	_, err := registerEmployee()
	if err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusCreated, "insert error")
	}

	return c.JSON(http.StatusCreated, "insert success")
}

func registerEmployee() (int64, error) {
	// DB コネクト
	db := core.DatabseConnect()

	// SQL 文用意
	sql := `INSERT INTO employees (emp_no, birth_date, first_name, last_name, gender, hire_date)
	VALUES (?, ?, ?, ?, ?, ?);`;

	// SQL 実行
	result, err := db.Exec(sql, 998, "2023-05-22", "example_fn", "example_ln", "M", "2023-05-22");
	if err != nil {
		return 0, fmt.Errorf("insert error : %v", err)
	}

	// 最後にインサートされた id を取得することでエラーかどうか判定するっぽい
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insert error: %v", err)
	}

	return id, nil
}