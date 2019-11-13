package db


import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"study/util"
)

// Сохранение информации в слайс
func F4() {
	dbConnector, err := sqlx.Open("postgres",
		"user=user password=1234567 dbname=phonebook sslmode=disable")
	util.FailOnError(err, "db connection")

	rows, err := dbConnector.Queryx("select * from person")
	util.FailOnError(err, "select query")

	for rows.Next() {
		slice, _ := rows.SliceScan()
		util.PrettyPrint(slice)
	}
	rows.Close()
}

