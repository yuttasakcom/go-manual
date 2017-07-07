package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := http.HandlerFunc(mux)
	err := http.ListenAndServe(":8080", h)
	fmt.Println(err)
}

// type indexHandler struct{}

// func (*indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello, Gopher."))
// }

func mux(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		indexHandler(w, r)
	case "/about":
		aboutHandler(w, r)
	default:
		//notFoundHandler(w, r)
		http.NotFound(w, r)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

// func notFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("404"))
// }
