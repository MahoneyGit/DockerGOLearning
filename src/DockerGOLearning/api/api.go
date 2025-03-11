package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MahoneyGit/DockerGOLearning.git/src/DockerGOLearning/logger"
	"github.com/MahoneyGit/DockerGOLearning.git/src/book"
	"github.com/go-playground/validator"
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

func getObjectById(writter http.ResponseWriter, request *http.Request) {
	writter.WriteHeader(http.StatusOK)
	writter.Header().Set("Content-Type", "application/json")
	// ':=' is the short declaration operator, it allows declaring and initialising variables in one step
	bookID := request.PathValue("bookID")
	responseMessage := ""

	bookIDAsInt, err := strconv.Atoi(bookID)
	if err != nil {
		responseMessage = fmt.Sprintf("Invalid Book ID: %s", bookID)
	} else {
		bookFoundByID, err := book.GetBookByID(bookIDAsInt)
		if err != nil {
			responseMessage = "Something went wrong, book not found. Please ensure the book has been created correctly"
		} else {
			responseMessage = bookFoundByID
		}
	}
	writter.Write([]byte(responseMessage))
}

func deleteObjectById(writter http.ResponseWriter, request *http.Request) {
	responseMessage := ""
	bookID := request.PathValue("bookID") // Todo how can I extract the convert logic out I wonder
	bookIDAsInt, err := strconv.Atoi(bookID)
	if err != nil {
		responseMessage = fmt.Sprintf("Invalid Book ID: %s", bookID)
	} else {
		bookDeleted, err := book.DeleteBookByID(bookIDAsInt)
		if err != nil || !bookDeleted {
			responseMessage = "Something went wrong, book not deleted, flushing database"
		} else if bookDeleted {
			responseMessage = fmt.Sprintf("Book ID: %s sucessfully deleted", bookID)
		}
	}
	writter.Write([]byte(responseMessage))
}

func createObject(writter http.ResponseWriter, request *http.Request) {
	// bookID := request.PathValue("bookID") //Todo get object details
	responseMessage := ""
	newBook, err := extractBookFromRequest(request)

	if err != nil {
		handleBadRequest(writter, fmt.Sprintf("Validation failed, Invalid Book: %v", err))
		return
	} else {
		fmt.Println("Great, new book is being created with the following details!")
		// Todo write to db
		// 	responseMessage = "Something went wrong, book not created. Check details and try again"
		newBook.PrintDetails()
		err := newBook.SaveBook()
		if err != nil {
			responseMessage = fmt.Sprintf("Something went horribly, oh so horribly wrong! %v", err)
		} else {
			responseMessage = fmt.Sprintf("Book ID: %d sucessfully created", newBook.BookID)
		}
	}
	writter.Write([]byte(responseMessage))
}

func updateObjectByID(writter http.ResponseWriter, request *http.Request) {
	responseMessage := ""
	updateBookRequest, err := extractBookFromRequest(request)

	if err != nil {
		handleBadRequest(writter, fmt.Sprintf("Validation failed, Invalid Request: %v", err))
		return
	}

	bookID := request.PathValue("bookID") // Todo how can I extract the convert logic out I wonder
	bookIDAsInt, err := strconv.Atoi(bookID)
	updateBookRequest.BookID = bookIDAsInt
	if err != nil {
		responseMessage = fmt.Sprintf("Invalid Book ID: %d", bookIDAsInt) //Todo can I extract this logic out sooner
		writter.Write([]byte(responseMessage))
	} else {
		fmt.Printf("You are attempting to update, %v, with the following details\n", updateBookRequest.BookID)
		updateBookRequest.PrintDetails()
		updateBookRequest.UpdateBook()
		responseMessage = fmt.Sprintf("Book ID: %d sucessfully updated!", updateBookRequest.BookID) //Todo except not really, integrate with db
		writter.Write([]byte(responseMessage))
	}
}

// s is a receiver variable, and *APIServer means that Run is a method on a pointer to an APIServer struct.
// This line is actually saying func | reciever type | function name | no variables being expected | returns an error type which can be nil
func (s *APIServer) Run() error {
	// ServeMux is something called a http request multplexer (router)
	router := http.NewServeMux()
	router.HandleFunc("GET /book/{bookID}", getObjectById)
	router.HandleFunc("PATCH /book/{bookID}", updateObjectByID)
	router.HandleFunc("PUT /book/create", createObject)
	router.HandleFunc("DELETE /book/{bookID}", deleteObjectById)

	middlewareChain := MiddlewareChain(
		logger.RequestLogger,
		requireAuth,
	)

	middlewareChain := MiddlewareChain(
		logger.RequestLogger,
		requireAuth,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
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

type Middleware func(http.Handler) http.HandlerFunc

// The ... before the type Middleware is a variadic parameter in Go. This means that the function MiddlewareChain can accept a variable number of arguments of type Middleware
// In other words, it allows you to pass any number of Middleware functions to MiddlewareChain, and these will be collected into a slice of Middleware inside the function.
func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next.ServeHTTP
	}
}

func handleBadRequest(writter http.ResponseWriter, responseBody string) {
	writter.WriteHeader(http.StatusBadRequest)
	writter.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = responseBody
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	writter.Write(jsonResp)
}

func extractBookFromRequest(request *http.Request) (book.Book, error) {
	//ToDo add as middleware for create/update
	validate := validator.New()
	var newBook book.Book

	decoder := json.NewDecoder(request.Body)
	decoder.Decode(&newBook) //Decoding the HTTP request body into the book variable:
	err := validate.Struct(newBook)
	if err != nil {
		fmt.Println("Error Occured")
		return book.Book{}, err // Todo, does this not just create a new object in memory? And is not ineffecient?
	}
	return newBook, nil
}
