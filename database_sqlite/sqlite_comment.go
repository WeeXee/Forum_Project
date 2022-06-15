package database_sqlite

import (
	"database/sql"
	"fmt"
	"log"
)

type Comment struct {
	IdComment int
	IdUser    string
	IdPost    int
	Comment   string
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
}

func CreateTableComment(db *sql.DB) {
	defer func(sqliteDatabase *sql.DB) {
		err3 := sqliteDatabase.Close()
		if err3 != nil {
			fmt.Println(err3)
		}
	}(db)
	createCommentTableSQL := `CREATE TABLE IF NOT EXISTS Comments(
    	IdComment INTEGER PRIMARY KEY AUTOINCREMENT,
    	"IdPost" INTEGER,
		"IdUser"  TEXT,
		"Comment" TEXT,                   
	  );` // SQL Statement for Create Table

	log.Println("Create Comment table...")
	statement, err := db.Prepare(createCommentTableSQL) // Prepare SQL Statement
	if err != nil {
		fmt.Println(err.Error())
	}
	_, _ = statement.Exec() // Execute SQL Statements
	log.Println("Comment table created")
}

func AddComment(comment Comment) {
	DoesFileExist("sqlite-database.db")
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
	log.Println("Inserting Comment record ...")
	insertCommentSQL := `INSERT INTO Comments( IdPost, IdUser, Comment) VALUES (?,?,?)`
	statement, err := db.Prepare(insertCommentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(comment.IdPost, comment.IdUser, comment.Comment)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println("Comment inserted !!")
	}
}

func GetComment() []Comment {
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		fmt.Println(err)
	}

	row, err := sqliteDatabase.Query("SELECT * FROM Comments ORDER BY IdPost")
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
		err := row.Scan(&idComment, &newComment.IdPost, &newComment.IdUser, &newComment.Comment)
		if err != nil {
			fmt.Println(err, "l'erreur est ici")
		}
		log.Println("IdPost :", newComment.IdPost, "IdUser: ", newComment.IdUser, "Comment :", newComment.Comment)
		commentArray = append(commentArray, newComment)
	}
	return commentArray
}
