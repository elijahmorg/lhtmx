package htmx

import (
	"strconv"

	"github.com/chasefleming/elem-go/htmx"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/styles"
)

// Todo model
type Todo struct {
	ID    int
	Title string
	Done  bool
}

// Global Todos slice (for simplicity)
var Todos = []Todo{
	{ID: 1, Title: "First task", Done: false},
	{ID: 2, Title: "Second task", Done: true},
}

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

func RenderTodos(todos []Todo) string {
	inputButtonStyle := styles.Props{
		styles.Width:           "100%",
		styles.Padding:         "10px",
		styles.MarginBottom:    "10px",
		styles.Border:          "1px solid #ccc",
		styles.BorderRadius:    "4px",
		styles.BackgroundColor: "#f9f9f9",
	}

	buttonStyle := styles.Props{
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

	listContainerStyle := styles.Props{
		styles.ListStyleType: "none",
		styles.Padding:       "0",
		styles.Width:         "100%",
	}

	centerContainerStyle := styles.Props{
		styles.MaxWidth:        "300px",
		styles.Margin:          "40px auto",
		styles.Padding:         "20px",
		styles.Border:          "1px solid #ccc",
		styles.BoxShadow:       "0px 0px 10px rgba(0,0,0,0.1)",
		styles.BackgroundColor: "#f9f9f9",
	}

	headContent := elem.Head(nil,
		elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org"}),
		elem.Script(attrs.Props{attrs.Src: "resources/wasm_exec.js"}),
		elem.Script(attrs.Props{attrs.Src: "resources/start_worker.js"}),
	)

	bodyContent := elem.Div(
		attrs.Props{attrs.Style: centerContainerStyle.ToInline()},
		elem.H1(nil, elem.Text("Todo List")),
		elem.Form(
			attrs.Props{attrs.Method: "post", attrs.Action: "/add"},
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

	htmlContent := elem.Html(nil, headContent, bodyContent)

	return htmlContent.Render()
}
