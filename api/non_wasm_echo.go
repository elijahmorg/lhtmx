//go:build !js

package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EchoStart() {
	fmt.Println("server side code")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(ServerDelay)

	// Routes
	e.GET("/", renderTodosRoute)
	e.POST("/toggle/:id", toggleTodoRoute)
	e.POST("/add", addTodoRoute)
	e.GET("/sync", getTodos)
	e.POST("/sync", syncTodos)

	e.Static("/", "../../public/")
	e.Logger.Fatal(e.Start(":3000"))
}
