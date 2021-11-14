package objects

import (
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates map[string]*template.Template
}

func NewTemplate(path string) *Template {
	templates := make(map[string]*template.Template)

	itens, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range itens {
		if !item.IsDir() && item.Name() != "base.html" && item.Name() != "header.html" && item.Name() != "footer.html" {
			templates[item.Name()] = template.Must(template.ParseFiles("static/"+item.Name(), "static/base.html", "static/header.html", "static/footer.html"))
		}
	}

	return &Template{
		templates: templates,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}
