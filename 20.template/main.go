package main

import (
	"html/template"
	"log"
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

type indexData struct {
	Name string
	List []string
}

var t = template.Must(template.ParseFiles("index.tmpl"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html;charset=utf-8")
	// t, err := template.ParseFiles("index.tmpl")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	data := indexData{
		Name: "Acoshift",
		List: []string{
			"Go",
			"C",
			"C++",
		},
	}

	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
