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
		fmt.Printf("Something went wrong! %v\n", err)
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Something went wrong! %v\n", err)
		panic(err)
	}

	fmt.Println("Connection Established")
}

func RunQuery(query string) (result *sql.Rows, err error) { // bad Practice to have just generic run query function but just testing
	fmt.Printf("Establishing environment and connection\n")
	establishEnvironment()
	fmt.Printf("Running query")

	dbQueryResult, err := db.Query(query)
	if err != nil {
		fmt.Printf("\n\nSomething has gone seriously wrong!%v\n\n", err)
		return nil, err
	}

	fmt.Println(dbQueryResult)
	return dbQueryResult, nil
}

func CloseConnection() {
	fmt.Println("Closing connection")
	db.Close()
}

func establishEnvironment() {
	EstablishConnection()
	createDevTables()
}

func createDevTables() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS public.book (
	book_id serial PRIMARY KEY,
	book_name VARCHAR(250) NOT NULL,
	num_of_pages int,
	book_description VARCHAR(250),
	publisher_id int)`
	fmt.Println(createTableQuery)
	// defer dbQueryResult.Close()
	// CloseConnection()
	dbQueryResult, err := db.Query(createTableQuery)
	if err != nil {
		fmt.Printf("Something went wrong! %v\n", err)
		panic(err)
	} else {
		fmt.Println(dbQueryResult)
	}
}
