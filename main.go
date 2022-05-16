package main

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, r)
}

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", Index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
