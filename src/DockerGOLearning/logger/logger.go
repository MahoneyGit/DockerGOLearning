package logger

import (
	"log"
	"net/http"
)

func RequestLogger(next http.Handler) http.HandlerFunc {
	return func(writter http.ResponseWriter, request *http.Request) {
		log.Printf("method %s, path: %s", request.Method, request.URL.Path)
		next.ServeHTTP(writter, request)
	}
}
