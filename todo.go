package main

import (
  "database/sql"
  "go-todo/handlers"

  "github.com/labstack/echo"
  _ "github.com/mattn/go-sqlite3"
)

func main() {
  db := initDB("storage.db")
  migrate(db)

  // Create a new instance of Echo
  e := echo.New()

  e.File("/", "public/index.html")
  e.GET("/tasks", handlers.GetTasks(db))
  e.PUT("/tasks", handlers.PutTask(db))
  e.DELETE("/tasks/:id", handlers.DeleteTask(db))

  // Start as a web server
  e.Start(":8000")
}

func initDB(filepath string) *sql.DB {
  db, err := sql.Open("sqlite3", filepath)

  // Check for any db errors then exit
  if err != nil {
    panic(err)
  }

  // If no errors but no db connection then exit
  if db == nil {
    panic("db nil")
  }
  return db
}

func migrate(db *sql.DB) {
  sql := `
  CREATE TABLE IF NOT EXISTS tasks(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name VARCHAR NOT NULL
  );
  `

  _, err := db.Exec(sql)
  // Exit if something is wrong with the SQL statement
  if err != nil {
    panic(err)
  }
}
