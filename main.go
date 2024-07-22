package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type IndexHandler struct{}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	fmt.Println("Starting server at 127.0.0.1:8080")

	h := IndexHandler{}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
