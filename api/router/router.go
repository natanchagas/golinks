package router

import (
	"github.com/labstack/echo/v4"
	"github.com/natanchagas/golinks/api"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//set main routes
	api.Endpoints(e)
	return e
}
