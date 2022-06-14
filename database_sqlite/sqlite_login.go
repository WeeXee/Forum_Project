package database_sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Login struct {
	Id       int
	Mail     string
	Username string
	Password string
}

type CommentLike struct {
	IdPost         int
	IdUser         int
	CommentContent string
	Like           int
	Dislike        int
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
		err2 := rows.Scan(&tempLogin.Id, &tempLogin.Mail, &tempLogin.Username, &tempLogin.Password)
		if err2 != nil {
			fmt.Println(err2)
		}
		if tempLogin.Mail == log.Mail && tempLogin.Password == log.Password {
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
    	"Mail" TEXT,
		"name" TEXT,
		"Password" TEXT		
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
	insertLoginSQL := `INSERT INTO login( Mail, name, Password) VALUES (?,?, ?)`
	statement, err := db.Prepare(insertLoginSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(login.Mail, login.Username, login.Password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

/*
func DisplayLogin(db *sql.DB) {
	row, err := db.Query("SELECT * FROM login ORDER BY Mail")
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
		var Password string
		err := row.Scan(&idLogin, &name, &Password)
		if err != nil {
			return
		}
		log.Println("Id :", idLogin, "login: ", name, " ", Password)
	}
}
*/

func checkIfLoginExist(db *sql.DB, login Login) bool {
	loginFree := true
	row, _ := db.Query("SELECT * FROM login ORDER BY Mail")
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
		if mail == login.Mail {
			loginFree = false
		}
	}
	return loginFree
}
