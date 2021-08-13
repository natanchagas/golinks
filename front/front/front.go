package front

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "home", template.HTML("<p>HTML Test</p>"))
}

func Endpoints(e *echo.Echo) {
	e.GET("/", Hello)
	// e.GET("/health", HealthCheck)
	// e.GET("/golink", CreateGolink)
	// e.GET("/golinks", ListGolinks)
	// e.GET("/:link", Golink)
}
