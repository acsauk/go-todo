package handlers

import (
  "database/sql"
  "net/http"
  "strconv"

  "go-todo/models"

  "github.com/labstack/echo"
)

// Convenience method to return JSON in responses
type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    // Fetch tasks using the GetTasks model
    return c.JSON(http.StatusOK, models.GetTasks(db))
  }
}

// PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    // Instantiate a new Task model
    var task models.Task
    // Map incoming JSON body to the new Task
    c.Bind(&task)
    // Add a task using the PutTask model
    id, err := models.PutTasks(db, task.name)
    // Return JSON response if successful
    if err == nil {
      return c.JSON(http.StatusCreated, H{
        "created": id,
      })
    // Handle errors
    } else {
      return err
    }
  }
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    // Use DeleteTask model to delete a task
    _, err := models.DeleteTask(db, id)
    // Return a JSON response on success
    if err == nil {
      return c.JSON(http.StatusOK, H{
        "deleted": id,
      })
    // Handle errors
    } else {
      return err
    }
  }
}
