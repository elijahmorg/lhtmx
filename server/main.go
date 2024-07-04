package main

import (
	"fmt"
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

	app.Static("/", "./public/")

	app.Listen(":3000")
}

func renderTodosRoute(c *fiber.Ctx) error {
	c.Type("html")
	return c.SendString(htmx.RenderTodos(htmx.Todos))
}

func toggleTodoRoute(c *fiber.Ctx) error {
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
		htmx.Todos = append(htmx.Todos, htmx.Todo{ID: len(htmx.Todos) + 1, Title: newTitle, Done: false})
	}
	return c.Redirect("/")
}
