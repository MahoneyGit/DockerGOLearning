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

func createDevTables() {
	createTableQuery := `CREATE TABLE IF NOT EXISTS devSchema.book (
	book_id SERIAL_PRIMARY_KEY,
	book_name VARCHAR(50) NOT NULL,
	num_of_pages int,
	book_description VARCHAR(150),
	publisher_id int)`
	fmt.Println(createTableQuery)
	// defer dbQueryResult.Close()
	CloseConnection()
}

// Todo I want to be able to create a connection and treat it as
// func (db *sql.DB) insertBook (bookName string, book)

// func updateBookTitle(title int, bookTitle string) {
// 	createTableQuery := "Update bookTitle FROM information_schema.tables where table_name = 'BasicTable'"

// 	defer dbQueryResult.Close()
// }
