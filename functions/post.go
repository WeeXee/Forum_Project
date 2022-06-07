package functions

import (
	"fmt"
	"net/http"
	"text/template"
)

type StructPost struct {
	MovieGender int
	IDuser      int
	post        string
	title       string
}

func Post(w http.ResponseWriter, r *http.Request) {
	newpost := StructPost{}
	newpost.title = r.FormValue("Title")
	newpost.post = r.FormValue("Containt")
	switch {
	case newpost == StructPost{}:
		break
	default:
		fmt.Println(newpost)
	}
	t, _ := template.ParseFiles("template/index.html")
	err := t.Execute(w, newpost)
	if err != nil {
		fmt.Print("error")
	}
}
