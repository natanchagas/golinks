package front

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/natanchagas/golinks/front/objects"
)

func MakeData(extra map[string]interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	data["API_ENDPOINT"] = os.Getenv("API_ENDPOINT")

	for key, value := range extra {
		data[key] = value
	}

	log.Println(data)

	return data
}

func Hello(c echo.Context) error {
	// return c.Render(http.StatusOK, "hello", template.HTML("<p>HTML Test</p>"))
	return c.Render(http.StatusOK, "home.html", MakeData(nil))
}

func ListGolinks(c echo.Context) error {

	resp, get_err := http.Get(os.Getenv("API_ENDPOINT") + "/golinks")
	if get_err != nil {
		fmt.Println(get_err)
	}

	body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		log.Fatalln(read_err)
	}

	var links []objects.Golink

	json.Unmarshal(body, &links)

	data := make(map[string]interface{})
	data["golinks"] = links

	return c.Render(http.StatusOK, "golinks.html", MakeData(data))
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
