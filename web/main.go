package main

import (
	"fmt"
	"github.com/elijahmorg/lhtmx/htmx"
)

func main() {
	fmt.Println("hello world I am here")
}

func renderTodosRoute() string {
	return htmx.RenderTodos(htmx.Todos)
}

func toggleTodoRoute(id int) string {
	var updatedTodo htmx.Todo
	for i, todo := range htmx.Todos {
		if todo.ID == id {
			htmx.Todos[i].Done = !todo.Done
			updatedTodo = htmx.Todos[i]
			break
		}
	}
	return htmx.CreateTodoNode(updatedTodo).Render()
}

func addTodoRoute(todoTitle string) string {
	if todoTitle != "" {
		htmx.Todos = append(htmx.Todos, htmx.Todo{ID: len(htmx.Todos) + 1, Title: todoTitle, Done: false})
	}
	return renderTodosRoute()
}
