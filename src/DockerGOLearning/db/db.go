package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres-tutorial"
)

var db *sql.DB

func EstablishConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Printf("Establishing connection to %s/%d at %s\n", host, port, time.Now().UTC())

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Established")
}

func RunQuery() {
	query := "SELECT table_name FROM information_schema.tables where table_name = 'BasicTable'"
	dbQueryResult, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer dbQueryResult.Close()
	for dbQueryResult.Next() {
		var (
			table_name string
		)
		if err := dbQueryResult.Scan(&table_name); err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", table_name)
	}
	if err := dbQueryResult.Err(); err != nil {
		panic(err)
	}
}

func CloseConnection() {
	fmt.Println("Closing connection")
	db.Close()
}
