package front

import (
	"net/http"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	// return c.Render(http.StatusOK, "hello", template.HTML("<p>HTML Test</p>"))
	return c.Render(http.StatusOK, "home.html", nil)
}

func Endpoints(e *echo.Echo) {
	e.Static("/static", "static")
	e.GET("/", Hello)
	// e.GET("/health", HealthCheck)
	// e.GET("/golink", CreateGolink)
	// e.GET("/golinks", ListGolinks)
	// e.GET("/:link", Golink)
}
