package database_sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Post struct {
	IdPost      int
	MailUser    string
	MovieGender []string
	PostTitle   string
	PostContent string
	Like        int
	Dislike     int
}

type PostsArray = []Post

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
}

func CreateTablePost(db *sql.DB) {
	createPostTableSQL := `CREATE TABLE IF NOT EXISTS Post(
    	idLogin INTEGER PRIMARY KEY AUTOINCREMENT,
		"IdUser"  TEXT,
		"MovieGender" TEXT,
		"PostTitle"   TEXT,
		"PostContent" TEXT,
		"Like"        INTEGER,
		"Dislike"     INTEGER		
	  );` // SQL Statement for Create Table

	log.Println("Create admin acess to Post...")
	statement, err := db.Prepare(createPostTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = statement.Exec() // Execute SQL Statements
	log.Println("Post table created")
}

func AddPost(post Post) {
	log.Println("Inserting post record ...")
	db, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		fmt.Println(err)
	}
	defer func(sqliteDatabase *sql.DB) {
		err3 := sqliteDatabase.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(db)
	insertLoginSQL := `INSERT INTO Post( IdUser, MovieGender, PostTitle, PostContent, Like, Dislike) VALUES (?,?,?,?,?,?)`
	statement, err := db.Prepare(insertLoginSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(post.MailUser, strings.Join(post.MovieGender, "/"), post.PostTitle, post.PostContent, post.Like, post.Dislike)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GetPost() PostsArray {
	DoesFileExist("sqlite-database.db")
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File

	row, err := sqliteDatabase.Query("SELECT * FROM Post ORDER BY IdUser")
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(row)
	postArray := PostsArray{}
	for row.Next() { // Iterate and fetch the records from result cursor
		post := Post{}
		var movieGender string
		err := row.Scan(&post.IdPost, &post.MailUser, &movieGender, &post.PostTitle, &post.PostContent, &post.Like, &post.Dislike)
		post.MovieGender = strings.Split(movieGender, "/")

		if err != nil {
			fmt.Println(err)
		}
		postArray = append(postArray, post)
	}
	return postArray
}
