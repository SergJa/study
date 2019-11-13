package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"study/util"
)

type Phonebook struct {
	Id int
	PersonName string `db:"name"`
	Phone string
}

// Возможности sqlx-пакета
func F5() {
	dbConnector, err := sqlx.Open("postgres",
		"user=user password=1234567 dbname=phonebook sslmode=disable")
	util.FailOnError(err, "db connection")

	// биндим параметры по именам
	insertPb := Phonebook{7, "Kolia","15473"}
	// второй вариант - использовать map[string]interface{}
	_, err = dbConnector.NamedExec(
		`insert into person values (:id, :name, :phone)`, insertPb)
	util.FailOnError(err, "insert")

	slice := []Phonebook{}
	fmt.Println(dbConnector.Select(&slice, "select * from person"))
	util.PrettyPrint(slice)

	rows, err := dbConnector.Queryx("select * from person")
	util.FailOnError(err, "select query")

	for rows.Next() {
		row := Phonebook{}
		fmt.Println(rows.StructScan(&row))
		util.PrettyPrint(row)
	}
	rows.Close()
}
