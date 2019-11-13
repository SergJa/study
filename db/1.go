package db

import (
	"database/sql"
	"fmt"
	"study/util"
	_ "github.com/lib/pq"
)

// Получение информации о колонках в результате запроса
func F1() {
	dbConnector, err := sql.Open("postgres",
		"user=user password=1234567 dbname=phonebook sslmode=disable")
	util.FailOnError(err, "db connection")

	rows, err := dbConnector.Query("select * from person")
	util.FailOnError(err, "select query")

	fmt.Print("Column names: ")
	fmt.Println(rows.Columns())

	fmt.Print("Column types: ")
	types, _ := rows.ColumnTypes()
	for _, Type := range types {
		fmt.Print(Type.ScanType(), " ")
	}
}
