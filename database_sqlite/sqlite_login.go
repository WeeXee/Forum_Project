package database_sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Login struct {
	id       int
	mail     string
	username string
	password string
}

type CommentLike struct {
	idPost         int
	idUser         int
	commentContent string
	like           int
	dislike        int
}

func DatabaseLogin(log Login) {
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
	CreateTableLogin(sqliteDatabase) // Create Database Tables*/

	// INSERT RECORDS
	boolean := checkIfLoginExist(sqliteDatabase, log)
	if boolean == true {
		AddLogin(sqliteDatabase, log)
	} else {

	}
	// DISPLAY INSERTED RECORDS

}

func GetLogin(db *sql.DB, log Login) Login {
	rows, err := db.Query("select * from login")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var tempLogin Login
		err2 := rows.Scan(&tempLogin.id, &tempLogin.mail, &tempLogin.username, &tempLogin.password)
		if err2 != nil {
			fmt.Println(err2)
		}
		if tempLogin.mail == log.mail && tempLogin.password == log.password {
			return tempLogin
		}
	}
	return Login{}
}

func DoesFileExist(fileName string) {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		OsCreateFile()
	} else {
		fmt.Printf("%v file exist\n", fileName)
	}
}

func OsCreateFile() {
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	err2 := file.Close()
	if err2 != nil {
		fmt.Println(err2)
	}
	log.Println("sqlite-database.db created")
}

func CreateTableLogin(db *sql.DB) {
	createLoginTableSQL := `CREATE TABLE IF NOT EXISTS login(
    	idLogin INTEGER PRIMARY KEY AUTOINCREMENT,
    	"mail" TEXT,
		"name" TEXT,
		"password" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create admin acess...")
	statement, err := db.Prepare(createLoginTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = statement.Exec() // Execute SQL Statements
	log.Println("Admin table created")
}

// AddLogin We are passing db reference connection from main to our method with other parameters
func AddLogin(db *sql.DB, login Login) {
	log.Println("Inserting login record ...")
	insertLoginSQL := `INSERT INTO login( mail, name, password) VALUES (?,?, ?)`
	statement, err := db.Prepare(insertLoginSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(login.mail, login.username, login.password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

/*
func DisplayLogin(db *sql.DB) {
	row, err := db.Query("SELECT * FROM login ORDER BY mail")
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(row)
	for row.Next() { // Iterate and fetch the records from result cursor
		var idLogin int
		var name string
		var password string
		err := row.Scan(&idLogin, &name, &password)
		if err != nil {
			return
		}
		log.Println("id :", idLogin, "login: ", name, " ", password)
	}
}
*/

func checkIfLoginExist(db *sql.DB, login Login) bool {
	loginFree := true
	row, _ := db.Query("SELECT * FROM login ORDER BY mail")
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(row)
	for row.Next() { // Iterate and fetch the records from result cursor
		var idLogin int
		var mail string
		var name string
		var password string
		_ = row.Scan(&idLogin, &mail, &name, &password)
		if mail == login.mail {
			loginFree = false
		}
	}
	return loginFree
}
