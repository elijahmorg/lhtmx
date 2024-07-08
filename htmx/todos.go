package htmx

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/chasefleming/elem-go/htmx"
	"github.com/google/uuid"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
)

// Todo model
type Todo struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Done   bool   `json:"done,omitempty"`
	TimeID int64  `json:"timestamp,omitempty"`
}

func getUUIDv7() uuid.UUID {
	u, _ := uuid.NewV7()
	return u
}

// Global Todos slice (for simplicity)
var Todos = []Todo{}

func CreateTodoNode(todo Todo) elem.Node {
	checkbox := elem.Input(attrs.Props{
		attrs.Type:    "checkbox",
		attrs.Checked: strconv.FormatBool(todo.Done),
		htmx.HXPost:   "/toggle/" + strconv.Itoa(todo.ID),
		htmx.HXTarget: "#todo-" + strconv.Itoa(todo.ID),
	})

	return elem.Li(attrs.Props{
		attrs.ID: "todo-" + strconv.Itoa(todo.ID),
	}, checkbox, elem.Span(attrs.Props{
		attrs.Style: styles.Props{
			styles.TextDecoration: elem.If(todo.Done, "line-through", "none"),
		}.ToInline(),
	}, elem.Text(todo.Title)))
}

var (
	inputButtonStyle = styles.Props{
		styles.Width:           "100%",
		styles.Padding:         "10px",
		styles.MarginBottom:    "10px",
		styles.Border:          "1px solid #ccc",
		styles.BorderRadius:    "4px",
		styles.BackgroundColor: "#f9f9f9",
	}

	buttonStyle = styles.Props{
		styles.BackgroundColor: "#007BFF",
		styles.Color:           "white",
		styles.BorderStyle:     "none",
		styles.BorderRadius:    "4px",
		styles.Cursor:          "pointer",
		styles.Width:           "100%",
		styles.Padding:         "8px 12px",
		styles.FontSize:        "14px",
		styles.Height:          "36px",
		styles.MarginRight:     "10px",
	}

	listContainerStyle = styles.Props{
		styles.ListStyleType: "none",
		styles.Padding:       "0",
		styles.Width:         "100%",
	}
	centerContainerStyle = styles.Props{
		styles.MaxWidth:        "300px",
		styles.Margin:          "40px auto",
		styles.Padding:         "20px",
		styles.Border:          "1px solid #ccc",
		styles.BoxShadow:       "0px 0px 10px rgba(0,0,0,0.1)",
		styles.BackgroundColor: "#f9f9f9",
	}
)

func RenderTodos(todos []Todo) string {
	if len(todos) == 0 {
		todos = []Todo{
			{ID: 0, Title: "Zero task", Done: false, TimeID: time.Now().UnixMicro()},
			{ID: 1, Title: "First task", Done: false, TimeID: time.Now().UnixMicro() + 1},
			{ID: 2, Title: "Second task", Done: true, TimeID: time.Now().UnixMicro() + 2},
		}
		Todos = todos
	}
	headContent := elem.Head(nil,
		elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org"}),
		elem.Script(attrs.Props{attrs.Src: "wasm_exec.js"}),
		elem.Script(attrs.Props{attrs.Src: "start_worker.js"}),
	)
	bodyContent := Body(todos)

	htmlContent := elem.Html(nil, headContent, bodyContent)

	return htmlContent.Render()
}

func RenderBody(todos []Todo) string {
	htmlContent := elem.Body(nil, Body(todos))
	return htmlContent.Render()
}

func Body(todos []Todo) *elem.Element {
	bodyContent := elem.Div(
		attrs.Props{attrs.Style: centerContainerStyle.ToInline()},
		elem.H1(nil, elem.Text("Todo List")),
		elem.Form(
			// attrs.Props{attrs.Method: "post", attrs.Action: "/add"},
			attrs.Props{
				htmx.HXPost:   "/add",
				htmx.HXTarget: "body",
				htmx.HXSwap:   "innerHTML",
			},
			elem.Input(
				attrs.Props{
					attrs.Type:        "text",
					attrs.Name:        "newTodo",
					attrs.Placeholder: "Add new task...",
					attrs.Style:       inputButtonStyle.ToInline(),
				},
			),
			elem.Button(
				attrs.Props{
					attrs.Type:  "submit",
					attrs.Style: buttonStyle.ToInline(),
				},
				elem.Text("Add"),
			),
		),
		elem.Ul(
			attrs.Props{attrs.Style: listContainerStyle.ToInline()},
			elem.TransformEach(todos, CreateTodoNode)...,
		),
	)
	return bodyContent
}

func MergeChanges(local, server []Todo) ([]Todo, error) {
	merged := make(map[string]Todo)

	for i, v := range local {
		fmt.Printf("local[%d] = %+v\n", i, v)
	}
	for i, v := range server {
		fmt.Printf("server[%d] = %+v\n", i, v)
	}

	// Add all local items to the merged map
	for _, item := range local {
		merged[item.Title] = item
	}

	// Merge server items, using the most recent version
	for _, serverItem := range server {
		if _, exists := merged[serverItem.Title]; !exists {
			merged[serverItem.Title] = serverItem
		}
	}

	// Convert map back to slice
	result := make([]Todo, 0, len(merged))
	for _, item := range merged {
		result = append(result, item)
	}
	slices.SortFunc(result, sortTodos)
	for i, v := range result {
		result[i].ID = i
		fmt.Printf("result[%d] = %+v\n", i, v)
	}

	return result, nil
}

func sortTodos(a Todo, b Todo) int {
	return int(a.TimeID - b.TimeID)
}
