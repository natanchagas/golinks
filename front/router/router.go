package router

import (
	"html/template"

	"github.com/labstack/echo"
	"github.com/natanchagas/golinks/front/front"
	"github.com/natanchagas/golinks/front/objects"
)

func New() *echo.Echo {

	t := &objects.Template{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}

	// create a new echo instance
	e := echo.New()

	e.Renderer = t

	//set main routes
	front.Endpoints(e)
	return e
}
