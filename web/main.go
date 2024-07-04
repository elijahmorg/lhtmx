package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/elijahmorg/lhtmx/htmx"
	"github.com/julienschmidt/httprouter"
	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	router := httprouter.New()
	router.GET("/", renderTodosRoute)
	router.POST("/toggle/:id", toggleTodoRoute)
	router.POST("/add", addTodoRoute)
	router.NotFound = notFound()
	// router.NotFound = notFound
	// router.GET("/", renderTodosRoute)
	// router.GET("/hello/:name", toggleTodoRoute)
	fmt.Println("hello world I am here")
	wasmhttp.Serve(router)
	select {}
}

func notFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("route not found")
		fmt.Println(r.URL)
		fmt.Fprintln(w, "This is the people handler.")
	})
}
func renderTodosRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("hello world I am here: renderTodosRoute")
	io.WriteString(w, htmx.RenderTodos(htmx.Todos))
	return
}

func toggleTodoRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("hello world I am here: toggleTodoRoute")
	var updatedTodo htmx.Todo
	id, _ := strconv.Atoi(ps.ByName("id"))
	for i, todo := range htmx.Todos {
		if todo.ID == id {
			htmx.Todos[i].Done = !todo.Done
			updatedTodo = htmx.Todos[i]
			break
		}
	}
	io.WriteString(w, htmx.CreateTodoNode(updatedTodo).Render())
}

func addTodoRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("hello world I am here: addTodoRoute")
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get a single value
	todoTitle := r.FormValue("newTodo")
	// todoTitle := ps.ByName("title")
	if todoTitle != "" {
		htmx.Todos = append(htmx.Todos, htmx.Todo{ID: len(htmx.Todos) + 1, Title: todoTitle, Done: false})
	}
	renderTodosRoute(w, r, ps)

	http.Redirect(w, r, "/", http.StatusFound)
}
