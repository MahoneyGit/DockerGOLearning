package main

import "github.com/MahoneyGit/DockerGOLearning.git/src/DockerGOLearning/api"

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
