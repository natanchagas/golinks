package front

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func MakeData(extra map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	data["API_ENDPOINT"] = os.Getenv("API_ENDPOINT")

	for key, value := range extra {
		fmt.Println(key, value)
		data[key] = value
	}

	return data
}

func Hello(c echo.Context) error {
	// return c.Render(http.StatusOK, "hello", template.HTML("<p>HTML Test</p>"))
	fmt.Println(MakeData(nil))
	return c.Render(http.StatusOK, "home.html", MakeData(nil))
}

func ListGolinks(c echo.Context) error {
	return c.Render(http.StatusOK, "golinks.html", MakeData(nil))
}

func HealthCheck(c echo.Context) error {
	return c.Render(http.StatusOK, "health.html", MakeData(nil))
}

func Endpoints(e *echo.Echo) {
	e.Static("/static", "static")
	e.GET("/", Hello)
	// e.GET("/golink", CreateGolink)
	e.GET("/golinks", ListGolinks)
	e.GET("/health", HealthCheck)
	// e.GET("/:link", Golink)
}
