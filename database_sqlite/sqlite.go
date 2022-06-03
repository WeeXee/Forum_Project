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
	username string
	password string
	//I created a struct with a struct to select the rows in the table and add data.
}

type Post struct {
	idPost      int
	idUser      int
	postTitle   string
	postContent string
	like        int
	dislike     int
}

type CommentLike struct {
	idPost         int
	idUser         int
	commentContent string
	like           int
	dislike        int
}

func Database() {
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
	CreateTable(sqliteDatabase) // Create Database Tables*/

	log := Login{
		username: "Juju",
		password: "SLB",
	}
	// INSERT RECORDS
	boolean := checkIfLoginExist(sqliteDatabase, log)
	if boolean == true {
		AddLogin(sqliteDatabase, "Juju", "SLB")
	} else {

	}

	login := GetLogin(sqliteDatabase, 3)
	fmt.Println(login)
	// DISPLAY INSERTED RECORDS
	DisplayLogin(sqliteDatabase)

}

func GetLogin(db *sql.DB, id2 int) Login {
	rows, err := db.Query("select * from login")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var tempLogin Login
		err2 := rows.Scan(&tempLogin.id, &tempLogin.username, &tempLogin.password)
		if err2 != nil {
			fmt.Println(err2)
		}
		if tempLogin.id == id2 {
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

func CreateTable(db *sql.DB) {
	createLoginTableSQL := `CREATE TABLE IF NOT EXISTS login(
    	idLogin INTEGER PRIMARY KEY AUTOINCREMENT,
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

// We are passing db reference connection from main to our method with other parameters
func AddLogin(db *sql.DB, name string, password string) {
	log.Println("Inserting login record ...")
	insertLoginSQL := `INSERT INTO login( name, password) VALUES (?, ?)`
	statement, err := db.Prepare(insertLoginSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(name, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func DisplayLogin(db *sql.DB) {
	row, err := db.Query("SELECT * FROM login ORDER BY name")
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

func checkIfLoginExist(db *sql.DB, login Login) bool {
	loginFree := true
	row, _ := db.Query("SELECT * FROM login ORDER BY name")
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
		_ = row.Scan(&idLogin, &name, &password)
		if name == login.username {
			loginFree = false
		}
	}
	return loginFree
}
