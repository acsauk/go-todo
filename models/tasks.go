package models

import (
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
)

// Task - a struct containing Task data
type Task struct {
  ID   int `json:"id"`
  Name String `json:"name"`
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
