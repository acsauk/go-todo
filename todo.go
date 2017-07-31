package main

import (
  "github.com/labstack/echo"
)

func main() {
  // Create a new instance of Echo
  e := echo.New()

  e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET tasks") })
  e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT tasks") })
  e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELTE Task "+c.Param("id")) })

  // Start as a web server
  e.Start(":8000")
}
