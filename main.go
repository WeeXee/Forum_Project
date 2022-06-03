package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type structure struct {
	username string
	post     string
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, r)
}

func main() {
	http.HandleFunc("/", post)
	fmt.Printf("Starting server got testing \n")
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", Index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "template/index.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err : %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website r.postfrom = %v\n", r.PostForm)
		username := r.FormValue("Your Name")
		post := r.FormValue("Your post")
		fmt.Fprintf(w, "Name = %s\n", username)
		fmt.Print(w, "Your post = %s\n", post)
	default:
		fmt.Fprintf(w, "Only get and Post")
	}
}

func word_limiter(s string, limit int) string {

	if strings.TrimSpace(s) == "" {
		return s
	}

	// convert string to slice
	strSlice := strings.Fields(s)

	// count the number of words
	numWords := len(strSlice)

	var result string

	if numWords > limit {
		// convert slice/array back to string
		result = strings.Join(strSlice[0:limit], " ")

		// the three dots for end characters are optional
		// you can change it to something else or remove this line
		result = result + "..."
	} else {

		// the number of limit is higher than the number of words
		// return default or else will cause
		// panic: runtime error: slice bounds out of range
		result = s
	}

	return string(result)

}
