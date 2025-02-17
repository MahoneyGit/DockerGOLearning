package api

import (
	"log"
	"net/http"

	"github.com/MahoneyGit/DockerGOLearning.git/src/DockerGOLearning/logger"
)

type APIServer struct {
	addr string
}

// * says return type is a pointer. THis means we are passing a reference to the underlying struct, this is more effecient than simply passing a copy of the struct
// the & then says to return the pointer of the object
func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func writeUser(writter http.ResponseWriter, request *http.Request) {
	bookID := request.PathValue("bookID")
	writter.Write([]byte("Book ID: " + bookID))
}

// s is a receiver variable, and *APIServer means that Run is a method on a pointer to an APIServer struct.
// This line is actually saying func | reciever type | function name | no variables being expected | returns an error type which can be nil
func (s *APIServer) Run() error {
	// ServeMux is something called a http request multplexer (router)
	router := http.NewServeMux()
	router.HandleFunc("GET /book/{bookID}", writeUser)

	server := http.Server{
		Addr:    s.addr,
		Handler: requireAuth(logger.RequestLogger(router)),
	}

	log.Printf("Server started on address %s", s.addr)

	// Server is started and ready to serve
	return server.ListenAndServe()
}

// Worth copying for a reminder explanation

// What is a Receiver in Go?
// A receiver allows you to define methods that operate on instances of a type (usually structs). The receiver is essentially the object (or struct) that the method is being invoked on. It is like the "this" keyword in other OOP languages like Java, C++, or Python.

// How Receivers Work in Go:
// You define the receiver when you define the method. It's specified between the func keyword and the method name, before the parameters.
// The receiver can be either a value receiver or a pointer receiver.
// Value reciever e.g.
// point := &Point{1, 2} // Pointer to a Point
// 	point.Move(5, 6) // This modifies the original point
// 	fmt.Println(*point) // Output: {6 8}

func requireAuth(next http.Handler) http.HandlerFunc {
	return func(writter http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		if token != "Bearer token" {
			http.Error(writter, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(writter, request)
	}
}
