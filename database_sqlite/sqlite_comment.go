package database_sqlite

import (
	"database/sql"
	"fmt"
	"log"
)

type Comment struct {
	idComment int
	idUser    int
	idPost    int
	comment   string
	like      int
	dislike   int
	answer    string
}

type CommentArray = []Comment

func DatabaseComment() {
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
	CreateTableComment(sqliteDatabase) // Create Database Tables*/

	comment := Comment{
		idUser:  1,
		idPost:  1,
		comment: "accord ave vows!",
		like:    10,
		dislike: 0,
		answer:  "blob || blab || blubber",
	}
	// INSERT RECORDS

	AddComment(sqliteDatabase, comment)
	// DISPLAY INSERTED RECORDS
}

func CreateTableComment(db *sql.DB) {
	createCommentTableSQL := `CREATE TABLE IF NOT EXISTS Comments(
    	idComment INTEGER PRIMARY KEY AUTOINCREMENT,
    	"idPost" INTEGER,
		"idUser"  INTEGER,
		"comment" TEXT,
		"like" INTEGER ,
		"dislike"  INTEGER,
        "Answer"  TEXT                      
	  );` // SQL Statement for Create Table

	log.Println("Create Comment table...")
	statement, err := db.Prepare(createCommentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = statement.Exec() // Execute SQL Statements
	log.Println("Comment table created")
}

func AddComment(db *sql.DB, comment Comment) {
	log.Println("Inserting comment record ...")
	insertCommentSQL := `INSERT INTO Comments( idPost, idUser, Comment, like, dislike, answer) VALUES (?,?,?,?,?,?)`
	statement, err := db.Prepare(insertCommentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(comment.idPost, comment.idUser, comment.comment, comment.like, comment.dislike, comment.answer)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GetComment() []Comment {
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		fmt.Println(err)
	}

	row, err := sqliteDatabase.Query("SELECT * FROM Comments ORDER BY idPost")
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(row)
	commentArray := CommentArray{}

	for row.Next() {
		newComment := Comment{}
		var idComment string // Iterate and fetch the records from result cursor
		err := row.Scan(&idComment, &newComment.idPost, &newComment.idUser, &newComment.comment, &newComment.like, &newComment.dislike, &newComment.answer)
		if err != nil {
			fmt.Println(err, "l'erreur est ici")
		}
		log.Println("idPost :", newComment.idPost, "idUser: ", newComment.idUser, "Comment :", newComment.comment, "like :", newComment.dislike)
		commentArray = append(commentArray, newComment)
	}
	return commentArray
}
