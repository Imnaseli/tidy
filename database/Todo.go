package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	logger "github.com/sijirama/tidy/Logger"
)

type Todo struct {
	Title       string
	Description string
}

func CreateTodoTable(db *sql.DB) {
	createTodoTableSQL := `CREATE TABLE todo (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"title" TEXT,
		"description" TEXT,
		"completed" BOOLEAN DEFAULT FALSE,
		"createdAt" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	  );` // SQL Statement for Create Table
	statement, err := db.Prepare(createTodoTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
		logger.LogToFile(fmt.Sprintf("Failed to create todo table: %v", err))
	}
	statement.Exec() // Execute SQL Statements
}

func InsertTodo(db *sql.DB, todo Todo) {
	log.Println("Inserting todo record ...")
	insertTodoSQL := `INSERT INTO todo(title, description) VALUES (?, ?)`
	statement, err := db.Prepare(insertTodoSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(todo.Title, todo.Description)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Your new todo has been created!!")
}

type TodoComplete struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

// DisplayTodos retrieves and prints all todos from the database
func DisplayTodos(db *sql.DB) []TodoComplete {
	rows, err := db.Query("SELECT * FROM todo ORDER BY createdAt DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []TodoComplete
	for rows.Next() {
		var t TodoComplete
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		if err == nil {
			todos = append(todos, t)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return todos
}
