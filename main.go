package main

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"unicode"

	"Forum/database_sqlite"
	_ "github.com/go-sql-driver/mysql"
	"time"

	_ "github.com/mattn/go-sqlite3"

	_ "Forum/functions"
	"fmt"

	_ "github.com/dgrijalva/jwt-go"

	"html/template"
	"net/http"
)

type logIndex struct {
	Username string
}

var jwtKey = []byte("my_secret_key")

// Create a struct to read the username and password from the request body
type Credentials struct {
	Mail     string `json:"mail"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	Mail     string `json:"mail"`
	jwt.StandardClaims
}

func cookies(c *http.Cookie, login logIndex) logIndex {
	if c != nil {
		tknStr := c.Value

		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println(http.StatusUnauthorized)
				return login
			}
			fmt.Println(http.StatusBadRequest)
			return login
		}
		if !tkn.Valid {
			fmt.Println(http.StatusUnauthorized)
			return login
		}
		login.Username = claims.Username
	}
	return login
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)
	NavBar(w, r)
	t, _ := template.ParseFiles("template/index.html")
	err1 := t.Execute(w, login)
	if err1 != nil {
		fmt.Print("error")
	}
}

func NavBarLogged(w http.ResponseWriter, r *http.Request) {
	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")

	login = cookies(c, login)
	t, _ := template.ParseFiles("template/navbar_logged.html")
	err1 := t.Execute(w, nil)
	if err1 != nil {
		fmt.Print("error")
	}
}

func NavBar(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/navbar.html")
	err1 := t.Execute(w, nil)
	if err1 != nil {
		fmt.Print("error")
	}
}

func PostLogged(w http.ResponseWriter, r *http.Request) {

	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)
	var post database_sqlite.Post
	post.PostTitle = r.FormValue("title")
	post.PostContent = r.FormValue("content")

	post.Like = 0
	post.Dislike = 0
	post.MovieGender = []string{
		r.FormValue("comedy"), r.FormValue("action"), r.FormValue("drama"), r.FormValue("fantasy"), r.FormValue("horror")}

	fmt.Println(post.MailUser)

	if post.MailUser != "" && post.PostTitle != "" && post.PostTitle != "" {
		database_sqlite.AddPost(post)
	}
	NavBarLogged(w, r)
	t, _ := template.ParseFiles("template/create_post.html")
	err1 := t.Execute(w, nil)
	if err1 != nil {
		fmt.Print("error")
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}

	t, _ := template.ParseFiles("template/index.html")
	err1 := t.Execute(w, login)
	if err1 != nil {
		fmt.Print("error")
	}
}

func Action(w http.ResponseWriter, r *http.Request) {
	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}

	t, _ := template.ParseFiles("template/action.html")
	err1 := t.Execute(w, login)
	if err1 != nil {
		fmt.Print("error")
	}
}

type GenderPage struct {
	LogCookies logIndex
	PostArray  database_sqlite.PostsArray
}

func Biobic(w http.ResponseWriter, r *http.Request) {
	var postArray = database_sqlite.GetPost()
	var arrayPosts database_sqlite.PostsArray
	for _, v := range postArray {
		for _, val := range v.MovieGender {
			if val == "biopic" {
				arrayPosts = append(arrayPosts, v)
			}
		}
	}

	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
		PostLogged(w, r)
	} else {
		NavBar(w, r)
	}

	/*genderPage := GenderPage{login, arrayPosts}*/

	t, _ := template.ParseFiles("template/biopic.html")
	err1 := t.Execute(w, login)
	if err1 != nil {
		fmt.Print("error")
	}
}

func Comedy(w http.ResponseWriter, r *http.Request) {
	commentArray := database_sqlite.GetComment()
	var postArray = database_sqlite.GetPost()
	var arrayPosts database_sqlite.PostsArray

	for _, v := range postArray {
		for _, val := range v.MovieGender {
			if val == "comedy" {
				arrayPosts = append(arrayPosts, v)
			}
		}
	}

	for _, v := range commentArray {
		for _, val := range arrayPosts {
			if v.IdPost == val.IdPost && v.Comment != "" {
				fmt.Println("v.IdPost = " + string(v.IdPost))
				fmt.Println("val.IdPost = " + string(val.IdPost))
				val.PostComment = append(val.PostComment, v)
				fmt.Println(val.PostComment)
			}
		}
	}
	for _, val := range arrayPosts {
		fmt.Println(val.IdPost)
		fmt.Println(val.PostComment)
	}

	login := logIndex{
		Username: "vous",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}

	var s database_sqlite.Comment
	s.IdUser = login.Username
	s.IdPost, _ = strconv.Atoi(r.FormValue("idPost"))
	s.Comment = r.FormValue("comment")

	if s.IdUser != "/" && s.IdPost != 0 && s.Comment != "" {
		database_sqlite.AddComment(s)
	}

	genderPage := GenderPage{login, arrayPosts}
	t, _ := template.ParseFiles("template/comedy.html")
	err1 := t.Execute(w, genderPage)
	if err1 != nil {
		fmt.Print("error")
	}
}

func Fantasy(w http.ResponseWriter, r *http.Request) {
	var postArray = database_sqlite.GetPost()
	var arrayPosts database_sqlite.PostsArray
	for _, v := range postArray {
		for _, val := range v.MovieGender {
			if val == "comedy" {
				arrayPosts = append(arrayPosts, v)
				fmt.Println(v)
			}
		}
	}

	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}
	genderPage := GenderPage{login, arrayPosts}
	t, _ := template.ParseFiles("template/fantasy.html")
	err1 := t.Execute(w, genderPage)
	if err1 != nil {
		fmt.Print("error")
	}
}

func Horror(w http.ResponseWriter, r *http.Request) {
	var postArray = database_sqlite.GetPost()
	var arrayPosts database_sqlite.PostsArray
	for _, v := range postArray {
		for _, val := range v.MovieGender {
			if val == "comedy" {
				arrayPosts = append(arrayPosts, v)
				fmt.Println(v)
			}
		}
	}

	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}
	genderPage := GenderPage{login, arrayPosts}
	t, _ := template.ParseFiles("template/horror.html")
	err1 := t.Execute(w, genderPage)
	if err1 != nil {
		fmt.Print("error")
	}
}

/**not done yet**/

func Drama(w http.ResponseWriter, r *http.Request) {
	var postArray = database_sqlite.GetPost()
	var arrayPosts database_sqlite.PostsArray
	for _, v := range postArray {
		for _, val := range v.MovieGender {
			if val == "comedy" {
				arrayPosts = append(arrayPosts, v)
				fmt.Println(v)
			}
		}
	}

	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}
	genderPage := GenderPage{login, arrayPosts}
	t, _ := template.ParseFiles("template/drama.html")
	err1 := t.Execute(w, genderPage)
	if err1 != nil {
		fmt.Print("error")
	}
}

func Romantic(w http.ResponseWriter, r *http.Request) {
	var postArray = database_sqlite.GetPost()
	var arrayPosts database_sqlite.PostsArray
	for _, v := range postArray {
		for _, val := range v.MovieGender {
			if val == "comedy" {
				arrayPosts = append(arrayPosts, v)
				fmt.Println(v)
			}
		}
	}

	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}
	genderPage := GenderPage{login, arrayPosts}

	t, _ := template.ParseFiles("template/romantic.html")
	err1 := t.Execute(w, genderPage)
	if err1 != nil {
		fmt.Print("error")
	}
}

func SF(w http.ResponseWriter, r *http.Request) {
	var postArray = database_sqlite.GetPost()
	var arrayPosts database_sqlite.PostsArray
	for _, v := range postArray {
		for _, val := range v.MovieGender {
			if val == "comedy" {
				arrayPosts = append(arrayPosts, v)
				fmt.Println(v)
			}
		}
	}

	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}
	genderPage := GenderPage{login, arrayPosts}
	t, _ := template.ParseFiles("template/SF.html")
	err1 := t.Execute(w, genderPage)
	if err1 != nil {
		fmt.Print("error")
	}
}
func Thriller(w http.ResponseWriter, r *http.Request) {
	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}

	t, _ := template.ParseFiles("template/thriller.html")
	err1 := t.Execute(w, login)
	if err1 != nil {
		fmt.Print("error")
	}
}

func Western(w http.ResponseWriter, r *http.Request) {
	var postArray = database_sqlite.GetPost()
	var arrayPosts database_sqlite.PostsArray
	for _, v := range postArray {
		for _, val := range v.MovieGender {
			if val == "comedy" {
				arrayPosts = append(arrayPosts, v)
				fmt.Println(v)
			}
		}
	}

	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	if login.Username != "/" {
		NavBarLogged(w, r)
	} else {
		NavBar(w, r)
	}
	genderPage := GenderPage{login, arrayPosts}
	t, _ := template.ParseFiles("template/western.html")
	err1 := t.Execute(w, genderPage)
	if err1 != nil {
		fmt.Print("/")
	}
}

func main() {
	database_sqlite.DatabaseCreateTable()
	http.HandleFunc("/", Index)
	http.HandleFunc("/action", Action)
	http.HandleFunc("/comedy", Comedy)
	http.HandleFunc("/biobic", Biobic)
	http.HandleFunc("/fantasy", Fantasy)
	http.HandleFunc("/horror", Horror)
	http.HandleFunc("/romantic", Romantic)
	http.HandleFunc("/SF", SF)
	http.HandleFunc("/thriller", Thriller)
	http.HandleFunc("/western", Western)

	/*formulaire Vanessa*/
	/***************************************************/

	http.HandleFunc("/getform", getFormHandler)
	/*http.HandleFunc("/processget", processGetHandler)*/
	http.HandleFunc("/postform", postFormHandler)
	http.HandleFunc("/processpost", processPostHandler)
	http.HandleFunc("/post", PostLogged)

	http.HandleFunc("/login", log)
	http.HandleFunc("/loginauth", Signin)

	http.HandleFunc("/logout", Logout)

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)

	/*Page note done*/
	http.HandleFunc("/drama", Drama)
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

	http.ListenAndServe("localhost:8080", nil)

}

var tpl *template.Template

func log(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func Signin(w http.ResponseWriter, r *http.Request) {

	fmt.Println("*****loginHandler running*****")
	var creds database_sqlite.Login

	creds.Mail = r.FormValue("mail")
	creds.Password = r.FormValue("password")
	creds.Username = r.FormValue("username")

	database_sqlite.DatabaseLogin(creds)
	expectedPassword, ok := database_sqlite.CheckLogin(creds)

	if !ok || expectedPassword != creds.Password {
		fmt.Println("error 2")
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		fmt.Println("succeed")
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		Mail:     creds.Mail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("error 3")
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	tpl.ExecuteTemplate(w, "login.html", nil)
}

// registerHandler serves form for registring new users
func registerHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("*****registerHandler running*****")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

// creates new user
func registerAuthHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("*****registerAuthHandler running*****")
	r.ParseForm()

	var user database_sqlite.Login
	//Create Username
	user.Mail = r.FormValue("mail")
	user.Username = r.FormValue("username")
	var nameAlphaNumeric = true
	for _, char := range user.Username {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}
	var nameLength bool
	if 4 <= len(user.Username) && len(user.Username) <= 20 {
		nameLength = true
	}
	// Create user.Password
	user.Password = r.FormValue("password")
	fmt.Println("user.Password:", user.Password, "\npswdLength:", len(user.Password))
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range user.Password {
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
	if 5 < len(user.Password) && len(user.Password) < 30 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {
		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
		return
	} else {
		added := database_sqlite.DatabaseLogin(user)
		if added {
			var creds Credentials
			creds.Mail = r.FormValue("mail")
			creds.Password = r.FormValue("password")
			creds.Username = r.FormValue("username")

			expirationTime := time.Now().Add(60 * time.Minute)
			claims := &Claims{
				Username: creds.Username,
				Mail:     creds.Mail,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Create the JWT string
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				fmt.Println("error 3")
				// If there is an error in creating the JWT return an internal server error
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Finally, we set the client cookie for "token" as the JWT we just generated
			// we also set an expiry time which is the same as the token itself
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})

			tpl.ExecuteTemplate(w, "register.html", "congrats, your account has been successfully created")
		} else {
			tpl.ExecuteTemplate(w, "register.html", "we meet a problem, retry please")
		}
	}
}

/******************formulaire**********************/

func getFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "getform.html", nil)
}

/*
func processGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processGetHandler running")
	login := logIndex{
		Username: "/",
	}
	c, _ := r.Cookie("token")
	login = cookies(c, login)

	var s database_sqlite.Comment
	s.IdUser = login.Username
	s.IdPost, _ = strconv.Atoi(r.FormValue("IdPost"))
	s.Comment = r.FormValue("textcomment")
	database_sqlite.AddComment(s)

	tpl.ExecuteTemplate(w, "action.html", s)
}*/

func postFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "postform.html", nil)
}

func processPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processPostHandler running")
	var postArray = database_sqlite.GetPost()

	var err error
	if err != nil {
		fmt.Printf("error parsing float64")
		fmt.Println("error 3")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "comedy.html", postArray)
}
