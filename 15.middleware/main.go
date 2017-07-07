package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := logger(http.HandlerFunc(indexHandler))
	err := http.ListenAndServe(":8080", h)
	fmt.Println(err)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s, path: %s, query: %s\n", r.RequestURI, r.URL.Path, r.URL.Query())
		h.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}
