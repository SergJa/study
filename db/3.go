package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"study/util"
)

// MapScan - сохранение значений в maps
func F3() {
	dbConnector, err := sqlx.Open("postgres",
		"user=user password=1234567 dbname=phonebook sslmode=disable")
	util.FailOnError(err, "db connection")

	rows, err := dbConnector.Queryx("select * from person")
	util.FailOnError(err, "select query")

	for rows.Next() {
		row := map[string]interface{}{}
		fmt.Println(rows.MapScan(row))
		util.PrettyPrint(row)
	}
	rows.Close()
}