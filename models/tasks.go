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
