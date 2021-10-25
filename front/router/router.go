package router

import (
	"github.com/labstack/echo/v4"
	"github.com/natanchagas/golinks/front/front"
	"github.com/natanchagas/golinks/front/objects"
)

func New() *echo.Echo {

	t := objects.NewTemplate("static/*.html")

	// create a new echo instance
	e := echo.New()

	e.Renderer = t

	//set main routes
	front.Endpoints(e)
	return e
}
