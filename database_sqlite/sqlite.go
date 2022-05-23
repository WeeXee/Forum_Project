package database_sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Database() {
	/*os.Remove("sqlite-database.db")*/

	doesFileExist("sqlite-database.db")
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
	createTable(sqliteDatabase) // Create Database Tables*/

	// INSERT RECORDS
	insertLogin(sqliteDatabase, "Tintin", "Milou")

	// DISPLAY INSERTED RECORDS
	displayLogin(sqliteDatabase)
}

func doesFileExist(fileName string) {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {

	} else {
		fmt.Printf("%v file exist\n", fileName)
	}
}

func osCreateFile() {
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

func createTable(db *sql.DB) {
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
func insertLogin(db *sql.DB, name string, password string) {
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

func displayLogin(db *sql.DB) {
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
		log.Println("login: ", name, " ", password)
	}
}
