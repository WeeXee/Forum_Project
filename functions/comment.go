package functions

import (
	"fmt"
	"net/http"
	"text/template"
)

type StructComment struct {
	comment string
}

func Comment(w http.ResponseWriter, r *http.Request) {
	newcomment := StructComment{}
	newcomment.comment = r.FormValue("comment")
	switch {
	case newcomment == StructComment{}:
		break
	default:
		fmt.Println(newcomment)
	}
	t, _ := template.ParseFiles("template/index.html")
	err := t.Execute(w, newcomment)
	if err != nil {
		fmt.Print("error")
	}
}
