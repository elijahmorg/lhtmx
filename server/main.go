package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/elijahmorg/lhtmx/htmx"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func main() {
	app := fiber.New()

	// Routes
	app.Get("/", renderTodosRoute)
	app.Post("/toggle/:id", toggleTodoRoute)
	app.Post("/add", addTodoRoute)
	app.Get("/sync", getTodos)
	app.Post("/sync", syncTodos)

	app.Static("/", "./public/")

	app.Listen(":3000")
}

func renderTodosRoute(c *fiber.Ctx) error {
	time.Sleep(1000 * time.Millisecond)
	c.Type("html")
	return c.SendString(htmx.RenderTodos(htmx.Todos))
}

func toggleTodoRoute(c *fiber.Ctx) error {
	time.Sleep(1000 * time.Millisecond)
	id, _ := strconv.Atoi(c.Params("id"))
	var updatedTodo htmx.Todo
	for i, todo := range htmx.Todos {
		if todo.ID == id {
			htmx.Todos[i].Done = !todo.Done
			updatedTodo = htmx.Todos[i]
			break
		}
	}
	c.Type("html")
	return c.SendString(htmx.CreateTodoNode(updatedTodo).Render())
}

func addTodoRoute(c *fiber.Ctx) error {
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("addTodoRouteServer")
	newTitle := utils.CopyString(c.FormValue("newTodo"))
	if newTitle != "" {
		htmx.Todos = append(htmx.Todos, htmx.Todo{ID: len(htmx.Todos) + 1, Title: newTitle, Done: false, TimeID: time.Now().Unix()})
	}
	return c.SendString(htmx.RenderBody(htmx.Todos))
}

func syncTodos(c *fiber.Ctx) error {
	time.Sleep(1000 * time.Millisecond)
	var todos []htmx.Todo
	err := c.BodyParser(&todos)
	if err != nil {
		return err
	}
	todos, _ = htmx.MergeChanges(todos, htmx.Todos)
	// fmt.Printf("htmx.Todos: %+v", htmx.Todos)
	htmx.Todos = todos
	// fmt.Printf("htmx.Todos: %+v", htmx.Todos)
	c.Status(http.StatusOK)
	c.JSON(htmx.Todos)
	fmt.Println("got new todos")

	return nil
}

func getTodos(c *fiber.Ctx) error {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("got new todos")

	return c.JSON(htmx.Todos)
}
