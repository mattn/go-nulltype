package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mattn/go-nulltype"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var nt nulltype.NullTime
	err = db.QueryRow("select current_timestamp").Scan(&nt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nt)
}
