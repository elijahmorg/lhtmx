package htmx

import (
	"html/template"
	"net/http"
)

type TmplTodo struct {
	ID      int
	Text    string
	Checked bool
}

type PageData struct {
	Todos []TmplTodo
}

func RenderIndex(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseGlob("../templates/*tmpl"))

	data := PageData{
		Todos: []TmplTodo{
			{ID: 1, Text: "First task", Checked: false},
			{ID: 2, Text: "Second task", Checked: true},
		},
	}
	tmpl.ExecuteTemplate(w, "base", data)
}
