package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", http.HandlerFunc(indexHandler))
	http.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("public")})))
	http.ListenAndServe(":8080", nil)
}

type noDir struct {
	http.Dir
}

func (d noDir) Open(name string) (http.File, error) {
	f, err := d.Dir.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!doctype html>
		<title>Static Web Server</title>
		<link href=/-/css/style.css rel=stylesheet>
		<p class=red>
			Static web server.
		</p>
	`))
}
