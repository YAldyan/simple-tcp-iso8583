package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type billing struct {
	idBilling string
	Nominal   int
	Status    string
}

func connect() (*sql.DB, error) {

	// user:password@tcp(host:port)/dbname
	// user@tcp(host:port)/dbname
	// user     => root
	// password => 
	// host     => 127.0.0.1 atau localhost
	// port     => 3306
	// dbname   => test
	db, err := sql.Open("mysql", "root:#@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQueryRow(id string) {
	var db, err = connect()

	checkError(err)

	defer db.Close()

	var result = billing{}

	err = db.
		QueryRow("select Nominal, Status from billing where idBilling = ?", id).Scan(&result.Nominal, &result.Status)

	checkError(err)

	fmt.Println("\nNominal :", result.Nominal, "\nStatus : ", result.Status)
}

func sqlUpdate(id string, status string) {
	db, err := connect()

	checkError(err)

	defer db.Close()

	_, err = db.Exec("update billing set Status = ? where idBilling = ?", status, id)

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	sqlQueryRow("1234567890123456")
	sqlUpdate("1234567890123456", "Y")
	sqlQueryRow("1234567890123456")
}
