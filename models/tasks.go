package models

import (
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
)

// Task - a struct containing Task data
type Task struct {
  ID   int `json:"id"`
  Name string `json:"name"`
}

// TaskCollection - a struct containing a collection of Tasks
type TaskCollection struct {
  Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
  sql := "SELECT * FROM tasks"
  rows, err := db.Query(sql)
  // Exit if the SQL doesn't work
  if err != nil {
    panic(err)
  }
  // Cleanup when the program exits
  defer rows.Close()

  result := TaskCollection{}
  for rows.Next() {
    task := Task{}
    err2 := rows.Scan(&task.ID, &task.Name)
    // Exit if we get an error
    if err2 != nil {
      panic(err2)
    }
    result.Tasks = append(result.Tasks, task)
  }
  return result
}

func GetTask(db *sql.DB, id int) Task {
  sql := "SELECT name FROM tasks WHERE id = ?"

  // Create a prepared SQL statement
  stmt, err := db.Prepare(sql)
  // Exit if the SQL doesn't work
  if err != nil {
    panic(err)
  }
  // Cleanup when the program exits
  defer stmt.Close()

  // Replace '?' in prepared statement with 'id'
  var name string
  err2 := stmt.QueryRow(id).Scan(&name)

  if err2 != nil {
    panic(err2)
  }

  return Task{ID: id, Name: name}
}

func PutTask(db *sql.DB, name string) (int64, error) {
  sql := "INSERT INTO tasks(name) VALUES(?)"

  // Create a prepared SQL statement
  stmt, err := db.Prepare(sql)
  // Exit if we get an error
  if err != nil {
    panic(err)
  }
  // Cleanup when program exits
  defer stmt.Close()

  // Replace '?' in prepared statement with 'name'
  result, err2 := stmt.Exec(name)
  // Exit if we get an error
  if err2 != nil {
    panic(err2)
  }

  return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
  sql := "DELETE FROM tasks WHERE id = ?"

  // Create a prepared SQL statement
  stmt, err := db.Prepare(sql)
  // Exit if we get an error
  if err != nil {
    panic(err)
  }

  // Replace '?' in prepared statement with 'id'
  result, err2 := stmt.Exec(id)
  // Exit if we get an error
  if err2 != nil {
    panic(err2)
  }

  return result.RowsAffected()
}
