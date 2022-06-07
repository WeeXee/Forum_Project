package main

import (
	"Forum/functions"
	"fmt"
	"Forum/database_sqlite"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, r)
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
	database_sqlite.DatabaseLogin()
	database_sqlite.DatabasePost()
	database_sqlite.DatabaseComment()
	arrayComment := database_sqlite.GetComment()
	fmt.Println(arrayComment)
	http.HandleFunc("/", functions.Post)
	http.HandleFunc("/1", functions.Login)
	fmt.Printf("Starting server got testing \n")
	fmt.Println("Go to this adress: localhost:8080")
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
