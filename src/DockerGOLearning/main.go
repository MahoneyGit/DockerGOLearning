package main

import (
	"github.com/MahoneyGit/DockerGOLearning.git/src/book"
)

func main() {
	// server := api.NewAPIServer(":8080")
	// server.Run()

	// db.EstablishConnection()
	// db.RunQuery()
	newBook, err := book.CreateBook("Manly Books", 350, "a great book, still being written", 27)
	if err != nil {
		panic(err)
	}
	newBook.PrintDetails()
	// db.CloseConnection()
}
