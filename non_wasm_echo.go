//go:build !js

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func echoStart() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", renderTodosRoute)
	e.POST("/toggle/:id", toggleTodoRoute)
	e.POST("/add", addTodoRoute)
	// e.Server.Handler
	e.Logger.Fatal(e.Start(":3000"))
}
