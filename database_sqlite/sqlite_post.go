package database_sqlite

import (
	"database/sql"
	"fmt"
	"log"
)

type Post struct {
	idPost      int
	idUser      int
	movieGender string
	postTitle   string
	postContent string
	postComment string
	like        int
	dislike     int
}

type postsArray = []Post

func DatabasePost() {
	DoesFileExist("sqlite-database.db")
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		fmt.Println(err)
	}
	defer func(sqliteDatabase *sql.DB) {
		err3 := sqliteDatabase.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(sqliteDatabase) // Defer Closing the database
	CreateTablePost(sqliteDatabase) // Create Database Tables*/

	movieGender := []int{1, 2}
	var movieGenderString string

	for _, value := range movieGender {
		movieGenderString += string(value)
	}

	var postComment = []string{"c'est top!", "entièrement d'accord!"}
	var postcommentstring string
	for _, value := range postComment {
		postcommentstring += value
	}

	post := Post{
		idUser:      1,
		movieGender: "blabla",
		postTitle:   "first article",
		postContent: "this the first post in the forum web site, congratulations!",
		postComment: postcommentstring,
		like:        10,
		dislike:     0,
	}
	// INSERT RECORDS

	AddPost(sqliteDatabase, post)

	// DISPLAY INSERTED RECORDS
	GetPost()
}

func CreateTablePost(db *sql.DB) {
	createPostTableSQL := `CREATE TABLE IF NOT EXISTS Post(
    	idLogin INTEGER PRIMARY KEY AUTOINCREMENT,
		"idUser"  INTEGER,
		"movieGender" TEXT,
		"postTitle"   TEXT,
		"postContent" TEXT,
		"postComment" TEXT,
		"like"        INTEGER,
		"dislike"     INTEGER		
	  );` // SQL Statement for Create Table

	log.Println("Create admin acess...")
	statement, err := db.Prepare(createPostTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = statement.Exec() // Execute SQL Statements
	log.Println("Admin table created")
}

func AddPost(db *sql.DB, post Post) {
	log.Println("Inserting post record ...")
	insertLoginSQL := `INSERT INTO Post( idUser, movieGender, postTitle, postContent, postComment, like, dislike) VALUES (?,?,?,?,?,?,?)`
	statement, err := db.Prepare(insertLoginSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(post.idUser, post.movieGender, post.postTitle, post.postContent, post.postComment, post.like, post.dislike)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GetPost() []Post {
	DoesFileExist("sqlite-database.db")
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File

	row, err := sqliteDatabase.Query("SELECT * FROM Post ORDER BY idUser")
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(row)
	postArray := postsArray{}
	for row.Next() { // Iterate and fetch the records from result cursor
		post := Post{}
		err := row.Scan(&post.idPost, &post.idUser, &post.movieGender, &post.postTitle, &post.postContent, &post.postComment, &post.like, &post.dislike)
		if err != nil {
			fmt.Println(err)
		}
		postArray = append(postArray, post)
	}
	return postArray
}
