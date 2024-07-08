//go:build js

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	wasmhttp "github.com/nlepage/go-wasm-http-server"
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

	// Start server
	wasmhttp.Serve(e.Server.Handler)
	select {}
}
