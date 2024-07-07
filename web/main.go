package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/elijahmorg/lhtmx/htmx"
	"github.com/labstack/echo/v4"
)

func main() {
	err := getData()
	if err != nil {
		fmt.Println("error syncing data with server")
	}
	go getData()

	echoStart()
}

func notFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("route not found")
		fmt.Println(r.URL)
		fmt.Fprintln(w, "This is the people handler.")
	})
}
func renderTodosRoute(c echo.Context) error {
	fmt.Println("hello world I am here: renderTodosRoute")
	c.HTML(http.StatusOK, htmx.RenderTodos(htmx.Todos))
	syncData()
	return nil
}

func toggleTodoRoute(c echo.Context) error {
	fmt.Println("hello world I am here: toggleTodoRoute")
	var updatedTodo htmx.Todo
	id, _ := strconv.Atoi(c.Param("id"))
	for i, todo := range htmx.Todos {
		if todo.ID == id {
			htmx.Todos[i].Done = !todo.Done
			updatedTodo = htmx.Todos[i]
			break
		}
	}
	err := c.HTML(http.StatusOK, htmx.CreateTodoNode(updatedTodo).Render())
	syncData()
	return err
}

func addTodoRoute(c echo.Context) error {
	fmt.Println("hello world I am here: addTodoRoute")
	todoTitle := c.FormValue("newTodo")
	fmt.Println("TodoTitle: ", todoTitle)
	if todoTitle == "" {
		return c.String(http.StatusBadRequest, "no title provided")
	}

	// Get a single value
	todo := htmx.Todo{ID: len(htmx.Todos) + 1, Title: todoTitle, Done: false, TimeID: time.Now().Unix()}
	if todoTitle != "" {
		htmx.Todos = append(htmx.Todos, todo)
	}
	fmt.Println("hello world I am here: writing response")
	err := c.HTML(http.StatusOK, htmx.RenderBody(htmx.Todos))

	syncData()
	return err
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
