package views

import (
	"html/template"
	"io"
)

var templates Templates

type Templates struct {
	templates *template.Template
}

func Render(w io.Writer, name string, data any) error {
	return templates.templates.ExecuteTemplate(w, name, data)
}

func Init() {
	templates = Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}
