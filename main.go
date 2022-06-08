package main

import (
	"Forum/functions"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	functions.Post(w, r)
}
func Comedy(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/comedy.html")
	t.Execute(w, r)
}

func Docu(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/docu.html")
	t.Execute(w, r)
}

func Drama(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/drama.html")
	t.Execute(w, r)
}

func Horror(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/horror.html")
	t.Execute(w, r)
}

func Romantic(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/romantic.html")
	t.Execute(w, r)
}

func SF(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/SF.html")
	t.Execute(w, r)
}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/1", functions.Login)
	http.HandleFunc("/comedy", Comedy)
	http.HandleFunc("/docu", Docu)
	http.HandleFunc("/drama", Drama)
	http.HandleFunc("/horror", Horror)
	http.HandleFunc("/romantic", Romantic)
	http.HandleFunc("/SF", SF)
	fmt.Printf("Starting server got testing \n")
	fmt.Println("Go to this adress: localhost:8080")
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
