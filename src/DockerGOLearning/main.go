package main

import (
	"github.com/MahoneyGit/DockerGOLearning.git/src/DockerGOLearning/api"
	"github.com/MahoneyGit/DockerGOLearning.git/src/DockerGOLearning/db"
	"github.com/MahoneyGit/DockerGOLearning.git/src/book"
)

func main() {
	runServer()
	// CreateBook()
	// runDatabase()

}

func runServer() {
	server := api.NewAPIServer(":8080")
	server.Run()
}

func createBook() {
	newBook, err := book.CreateBook("Manly Books", 350, "a great book, still being written", 27)
	if err != nil {
		panic(err)
	}
	newBook.PrintDetails()
}

func runDatabase() {
	db.EstablishConnection()
	db.RunQuery()

	db.CloseConnection()
}
