package functions

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
		name := r.FormValue("name")
		firstname := r.FormValue("firstname")
		id := r.FormValue("id")
		password := r.FormValue("password")

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Firstname = %s\n", firstname)
		fmt.Fprintf(w, "ID = %s\n", id)
		fmt.Fprintf(w, "Password = %s\n", password)
		fmt.Fprintf(w, "Password = %s\n", password)

	default:
		fmt.Fprintf(w, "Only get and Post")
	}
}
