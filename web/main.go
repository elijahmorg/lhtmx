package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/elijahmorg/lhtmx/htmx"
	"github.com/julienschmidt/httprouter"
	wasmhttp "github.com/nlepage/go-wasm-http-server"
)

func main() {
	err := getData()
	if err != nil {
		fmt.Println("error syncing data with server")
	}

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
	getData()
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
	syncData()
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
	syncData()
}

func addTodoRoute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("hello world I am here: addTodoRoute")
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		fmt.Println("hello world I am here: error: ", err)
		return
	}

	// Get a single value
	todoTitle := r.FormValue("newTodo")
	// todoTitle := ps.ByName("title")
	todo := htmx.Todo{ID: len(htmx.Todos) + 1, Title: todoTitle, Done: false, TimeID: time.Now().Unix()}
	if todoTitle != "" {
		htmx.Todos = append(htmx.Todos, todo)
	}
	// renderTodosRoute(w, r, ps)
	fmt.Println("hello world I am here: writing response")
	io.WriteString(w, htmx.RenderBody(htmx.Todos))

	syncData()
}

func syncData() {
	go syncDataRoutine()
}

func syncDataRoutine() {
	b := bytes.NewBuffer([]byte(""))
	json.NewEncoder(b).Encode(htmx.Todos)

	resp, err := http.Post("http://localhost:3000/sync", "application/json", b)
	if err != nil {
		fmt.Println("error syncing data: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New("bad status code for sync")
		fmt.Println("error syncing data: ", err)
		return
	}
	todos := make([]htmx.Todo, 0)
	err = json.NewDecoder(resp.Body).Decode(&todos)
	if err != nil {
		fmt.Println("error decoding response: ", err)
	}

	todos, err = htmx.MergeChanges(htmx.Todos, todos)
	if err != nil {
		fmt.Println("error merging: ", err)
	}

	htmx.Todos = todos
}

func getData() error {

	fmt.Println("get data from server for syncing")
	resp, err := http.Get("http://localhost:3000/sync")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New("bad status code for sync")
		fmt.Println(err)
		return err
	}
	todos := make([]htmx.Todo, 0)
	json.NewDecoder(resp.Body).Decode(&todos)

	todos, err = htmx.MergeChanges(htmx.Todos, todos)
	if err != nil {
		return err
	}

	htmx.Todos = todos
	return nil
}
