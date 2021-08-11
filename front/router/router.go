package router

import (
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//set main routes
	front.Endpoints(e)
	return e
}
