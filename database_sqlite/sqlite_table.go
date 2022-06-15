package database_sqlite

import (
	"database/sql"
	"fmt"
	"log"
)

func DatabaseCreateTable() {
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
	createLoginTableSQL := `CREATE TABLE IF NOT EXISTS login(
    	idLogin INTEGER PRIMARY KEY AUTOINCREMENT,
    	"Mail" TEXT,
		"Name" TEXT,
		"Password" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create admin acess...")
	statement, err := db.Prepare(createLoginTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = statement.Exec() // Execute SQL Statements
	log.Println("Admin table created")

	createCommentTableSQL := `CREATE TABLE IF NOT EXISTS Comments(
	    	IdComment INTEGER PRIMARY KEY AUTOINCREMENT,
	    	"IdPost" INTEGER,
			"IdUser"  TEXT,
			"Comment" TEXT
		  );` // SQL Statement for Create Table

	log.Println("Create Comment table...")
	statement1, err1 := db.Prepare(createCommentTableSQL) // Prepare SQL Statement
	if err1 != nil {
		fmt.Println(err.Error())
	}

	_, _ = statement1.Exec() // Execute SQL Statements
	log.Println("Comment table created")

	createPostTableSQL := `CREATE TABLE IF NOT EXISTS Post(
    	idLogin INTEGER PRIMARY KEY AUTOINCREMENT,
		"IdUser"  TEXT,
		"MovieGender" TEXT,
		"PostTitle"   TEXT,
		"PostContent" TEXT,
		"PostComment" TEXT,
		"Like"        INTEGER,
		"Dislike"     INTEGER		
	  );` // SQL Statement for Create Table

	log.Println("Create admin acess to Post...")
	statement2, err3 := db.Prepare(createPostTableSQL) // Prepare SQL Statement
	if err3 != nil {
		log.Fatal(err.Error())
	}
	_, _ = statement2.Exec() // Execute SQL Statements
	log.Println("Post table created")

}
