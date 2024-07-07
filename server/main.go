package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/elijahmorg/lhtmx/htmx"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	echoStart()
}

func echoStart() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", renderTodosRoute)
	e.POST("/toggle/:id", toggleTodoRoute)
	e.POST("/add", addTodoRoute)
	e.GET("/sync", getTodos)
	e.POST("/sync", syncTodos)

	e.Static("/", "./public/")
	e.Logger.Fatal(e.Start(":3000"))
}

func renderTodosRoute(c echo.Context) error {
	time.Sleep(1000 * time.Millisecond)
	return c.HTML(http.StatusOK, htmx.RenderTodos(htmx.Todos))
}

func toggleTodoRoute(c echo.Context) error {
	time.Sleep(1000 * time.Millisecond)
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTodo htmx.Todo
	for i, todo := range htmx.Todos {
		if todo.ID == id {
			htmx.Todos[i].Done = !todo.Done
			updatedTodo = htmx.Todos[i]
			break
		}
	}
	return c.HTML(http.StatusOK, htmx.CreateTodoNode(updatedTodo).Render())
}

func addTodoRoute(c echo.Context) error {
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("addTodoRouteServer")
	// newTitle := utils.CopyString(c.FormValue("newTodo"))
	newTitle := c.FormValue("newTodo")
	if newTitle != "" {
		htmx.Todos = append(htmx.Todos, htmx.Todo{ID: len(htmx.Todos) + 1, Title: newTitle, Done: false, TimeID: time.Now().Unix()})
	}
	return c.HTML(http.StatusOK, htmx.RenderBody(htmx.Todos))
}

func syncTodos(c echo.Context) error {
	time.Sleep(1000 * time.Millisecond)
	var todos []htmx.Todo
	err := c.Bind(&todos)
	if err != nil {
		return err
	}
	todos, _ = htmx.MergeChanges(todos, htmx.Todos)
	htmx.Todos = todos
	c.JSON(http.StatusOK, htmx.Todos)
	fmt.Println("got new todos")

	return nil
}

func getTodos(c echo.Context) error {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("got new todos")

	return c.JSON(http.StatusOK, htmx.Todos)
}
