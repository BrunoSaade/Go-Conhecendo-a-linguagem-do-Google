package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:brumasa8@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	smt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	smt.Exec(2000, "Dagmar")
	smt.Exec(3000, "Rafael")
	_, err = smt.Exec(3000, "Giovana")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
