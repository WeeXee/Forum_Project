package functions

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type StructPost struct {
	MovieGender int
	IDuser      int
	post        string
	title       string
	like        int
	dislike     int
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

func WordLimiter(s string, limit int) string {

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
