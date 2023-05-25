package transactiondb

import (
	"fmt"
	_"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"echo/core"
)

// curl -X POST http://localhost/transactiondb で試せる
func Index(c echo.Context) error {
	// _, err := registerEmployeeInTransaction()
	_, err := registerEmployeeInTransactionPatterSuccess()
	if err != nil {
		fmt.Printf("registerEmployeeInTransaction でエラー発生")
		return c.JSON(http.StatusCreated, "insert in transaction error")
	}

	return c.JSON(http.StatusCreated, "insert in transaction success")
}

/*
* トランザクション内で同じテーブルに INSERT してみる
* 1つの INSERT は失敗する前提
*/
func registerEmployeeInTransactionPatterFail() (int64, error) {
	// DB コネクト
	db := core.DatabseConnect()

	// トランザクション開始
	tx, err := db.Begin()

	if err != nil {
		fmt.Println("db.Begin() FAIL!!")
	}

	// 失敗時にロールバックさせる
	defer tx.Rollback()

	fmt.Println("db.Begin() SUCCESS!!")

	// SQL 文用意
	sql := `INSERT INTO employees (emp_no, birth_date, first_name, last_name, gender, hire_date)
	VALUES (?, ?, ?, ?, ?, ?);`;

	// 成功するパターンの SQL 実行
	result, err := tx.Exec(sql, 996, "2023-05-25", "example_fn", "example_ln", "M", "2023-05-25");
	if err != nil {
		return 0, fmt.Errorf("insert error : %v", err)
	}

	// 失敗するパターンの SQL 実行
	// emp_no が 999 のレコードはすでに存在する前提です
	result, err = tx.Exec(sql, 999, "2023-05-25", "example_fn", "example_ln", "M", "2023-05-25");
	if err != nil {
		return 0, fmt.Errorf("insert error : %v", err)
	}

	// 最後にインサートされた id を取得することでエラーかどうか判定するっぽい
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("insert error: %v", err)
	}

	tx.Commit()
	fmt.Println("tx.Commit()")

	return id, nil
}

/*
* トランザクション内で異なるテーブルに INSERT してみる
* 2つの INSERT は成功する前提
*/
func registerEmployeeInTransactionPatterSuccess() (int64, error) {
	// DB コネクト
	db := core.DatabseConnect()

	// トランザクション開始
	tx, err := db.Begin()

	if err != nil {
		fmt.Println("db.Begin() FAIL!!")
	}

	// 失敗時にロールバックさせる
	defer tx.Rollback()

	fmt.Println("db.Begin() SUCCESS!!")

	// employees テーブルに INSERT
	sql_1 := `INSERT INTO employees (emp_no, birth_date, first_name, last_name, gender, hire_date)
	VALUES (?, ?, ?, ?, ?, ?);`;

	// 成功する想定
	result_1, err_1 := tx.Exec(sql_1, 990, "2023-05-25", "example_fn", "example_ln", "M", "2023-05-25");
	if err_1 != nil {
		return 0, fmt.Errorf("insert error : %v", err_1)
	}

	// 最後にインサートされた id を取得することでエラーかどうか判定する
	insert_id_1, err_insertid_1 := result_1.LastInsertId()
	if err_insertid_1 != nil {
		return 0, fmt.Errorf("insert error: %v", err_insertid_1)
	}
	fmt.Println("result_1.LastInsertId(): insert_id_1 => %v", insert_id_1)


	// salaries テーブルに INSERT
	sql_2 := `INSERT INTO salaries (emp_no, salary, from_date, to_date)
	VALUES (?, ?, ?, ?);`

	// 成功する想定
	result_2, err_2 := tx.Exec(sql_2, 990, 400000, "2023-05-25", "2025-01-01");
	if err_2 != nil {
		return 0, fmt.Errorf("insert error : %v", err_2)
	}

	// 最後にインサートされた id を取得することでエラーかどうか判定する
	insert_id_2, err_insertid_2 := result_2.LastInsertId()
	if err_insertid_2 != nil {
		return 0, fmt.Errorf("insert error: %v", err_insertid_2)
	}
	fmt.Println("result_1.LastInsertId(): insert_id_2 => %v", insert_id_2)


	tx.Commit()
	fmt.Println("tx.Commit()")

	return 1, nil
}