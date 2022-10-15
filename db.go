package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "104.198.73.247"
	port     = 5432
	user     = "postgres"
	password = "nD7zrq1e5r345t9u"
	dbname   = "postgres"
)

func main() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	// Delete
	deleteStmt := "delete from distributors where did in ($1,$2)"
	_, err = db.Exec(deleteStmt, 1, 2)
	CheckError(err)

	// dynamic
	insertDynStmt := `insert into distributors (name, did) values($1, $2)`
	_, err = db.Exec(insertDynStmt, "John", 1)
	CheckError(err)
	_, err = db.Exec(insertDynStmt, "Jane", 2)
	CheckError(err)

	// update
	updateStmt := `update distributors set name=$1 where did=$2`
	_, err = db.Exec(updateStmt, "Mary", 2)
	CheckError(err)

	rows, err := db.Query(`SELECT name, did FROM distributors`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var did int

		err = rows.Scan(&name, &did)
		CheckError(err)

		fmt.Println(name, did)
	}

	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
