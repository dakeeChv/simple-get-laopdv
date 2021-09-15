package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Conn() *sql.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Env("host"), Env("port"), Env("user"), Env("password"), Env("dbname"))

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
	return db
}
