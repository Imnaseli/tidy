package database

import (
	"database/sql"
	"fmt"
	"log"

	logger "github.com/sijirama/tidy/Logger"
)

func CreateTodoTable(db *sql.DB) {
	createTodoTableSQL := `CREATE TABLE todo (
		"idTodo" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"completed" BOOLEAN,
		"createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	  );` // SQL Statement for Create Table
	statement, err := db.Prepare(createTodoTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
        logger.LogToFile(fmt.Sprintf("Failed to create todo table: %v" , err))
	}
	statement.Exec() // Execute SQL Statements
}
