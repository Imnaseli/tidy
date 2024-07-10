package helper

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	logger "github.com/sijirama/tidy/Logger"
	"github.com/sijirama/tidy/database"
)

var DatabaseClient *sql.DB

var dbPathUrl = "db/databse.db"

func DatabseInit() {

	_, err := os.Stat(dbPathUrl)
	if err == nil {
		// Database file exists, open it
		sqliteDatabase, err := sql.Open("sqlite3", dbPathUrl)
		if err != nil {
			log.Fatal(err)
		}
		DatabaseClient = sqliteDatabase       // let other packages use the client connection
		logger.LogToFile("Database is ready") // log success
		return
	}

	file, err := os.Create(dbPathUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	sqliteDatabase, error := sql.Open("sqlite3", dbPathUrl)
	if error != nil {
		log.Fatal(error)
	}

	DatabaseClient = sqliteDatabase // let other packages use the client connection

	database.CreateTodoTable(DatabaseClient) // create some initial database tables

	logger.LogToFile("Database is ready") // log success

	//defer sqliteDatabase.Close()
}
