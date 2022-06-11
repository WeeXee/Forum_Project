package main

import (
	_ "github.com/go-sql-driver/mysql"
	"unicode"

	_ "github.com/mattn/go-sqlite3"

	"Forum/database_sqlite"
	"Forum/functions"
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, r)
}
func Action(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/action.html")
	t.Execute(w, r)
}

func Biobic(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/biobic.html")
	t.Execute(w, r)
}

func Comedy(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/comedy.html")
	t.Execute(w, r)
}

func Fantasy(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/fantasy.html")
	t.Execute(w, r)
}

func Horror(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/horror.html")
	t.Execute(w, r)
}

/**not done yet**/

func Drama(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/drama.html")
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
	http.HandleFunc("/action", Action)
	http.HandleFunc("/comedy", Comedy)
	http.HandleFunc("/biobic", Biobic)
	http.HandleFunc("/fantasy", Fantasy)
	http.HandleFunc("/horror", Horror)

	/*Page note done*/
	http.HandleFunc("/drama", Drama)
	http.HandleFunc("/romantic", Romantic)
	http.HandleFunc("/SF", SF)
	fmt.Printf("Starting server got testing \n")
	fmt.Println("Go to this adress: localhost:8080")
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))

	tpl, _ = template.ParseGlob("template/*.html")
	var err error
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.ListenAndServe("localhost:8080", nil)
}

//*vanessa partie*//

var tpl *template.Template

// registerHandler serves form for registring new users
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****registerHandler running*****")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

// creates new user
func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****registerAuthHandler running*****")
	r.ParseForm()
	//Create Username
	username := r.FormValue("username")
	var nameAlphaNumeric = true
	for _, char := range username {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
	var nameLength bool
	if 4 <= len(username) && len(username) <= 20 {
		nameLength = true
	}
	// Create password
	password := r.FormValue("password")
	fmt.Println("password:", password, "\npswdLength:", len(password))
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			pswdLowercase = true
		case unicode.IsUpper(char):
			pswdUppercase = true
		case unicode.IsNumber(char):
			pswdNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 5 < len(password) && len(password) < 30 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {
		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
		return
	}
	tpl.ExecuteTemplate(w, "register.html", "congrats, your account has been successfully created")
}
