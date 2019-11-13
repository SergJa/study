package db

import (
	"database/sql"
	"fmt"
	"study/util"
	_ "github.com/lib/pq"
)

// Добавление и чтение записей
func F2() {
	dbConnector, err := sql.Open("postgres",
		"user=user password=1234567 dbname=phonebook sslmode=disable")
	util.FailOnError(err, "db connection")

	name := "Вася Пупкин"
	id := 5

	_, err = dbConnector.Exec(
		`insert into person values ($2, $1, '8-123-456-7890')`,
		name, id)
	util.FailOnError(err, "insert")

	rows, err := dbConnector.Query("select * from person")
	util.FailOnError(err, "select query")

	for rows.Next() {
		var id int
		var name, phone string
		fmt.Println(rows.Scan(&id, &name, &phone))
		fmt.Println(id, name, phone)
	}
	rows.Close()
}
