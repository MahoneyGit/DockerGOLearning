package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Starting server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You have hit: %s\n", r.URL.Path)
		// http.Redirect(w, r, "http://www.google.com", http.StatusSeeOther)
	})

	http.ListenAndServe(":80", nil)
}
