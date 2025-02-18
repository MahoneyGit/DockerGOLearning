package main

import (
	"github.com/MahoneyGit/DockerGOLearning.git/src/DockerGOLearning/db"
)

func main() {
	// server := api.NewAPIServer(":8080")
	// server.Run()

	db.EstablishConnection()
	db.RunQuery()
	db.CloseConnection()
}
