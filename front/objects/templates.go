package objects

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func NewTemplate(path string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob(path)),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
