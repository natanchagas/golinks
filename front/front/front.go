package front

import (
	"bytes"
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

	return data
}

func Hello(c echo.Context) error {
	// return c.Render(http.StatusOK, "hello", template.HTML("<p>HTML Test</p>"))
	return c.Render(http.StatusOK, "home.html", MakeData(nil))
}

func CreateGolink(c echo.Context) error {

	link := c.FormValue("url")

	var g objects.Golink
	g.OriginalUrl = link

	json_data, json_err := json.Marshal(g)
	if json_err != nil {
		log.Println(json_err)
		return c.Render(http.StatusOK, "error.html", MakeData(nil))
	}

	resp, get_err := http.Post(os.Getenv("API_ENDPOINT")+"/golink", "application/json", bytes.NewBuffer(json_data))
	if get_err != nil {
		log.Println(get_err)
		return c.Render(http.StatusOK, "error.html", MakeData(nil))
	}

	var res objects.Golink
	json.NewDecoder(resp.Body).Decode(&res)

	data := make(map[string]interface{})
	data["golink"] = res

	return c.Render(http.StatusOK, "golink.html", MakeData(data))
}

func ListGolinks(c echo.Context) error {

	resp, get_err := http.Get(os.Getenv("API_ENDPOINT") + "/golinks")
	if get_err != nil {
		fmt.Println(get_err)
		return c.Render(http.StatusOK, "error.html", MakeData(nil))
	}

	body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		log.Fatalln(read_err)
		return c.Render(http.StatusOK, "error.html", MakeData(nil))
	}

	var links []objects.Golink

	json.Unmarshal(body, &links)

	data := make(map[string]interface{})
	data["golinks"] = links

	return c.Render(http.StatusOK, "golinks.html", MakeData(data))
}

func HealthCheck(c echo.Context) error {

	resp, get_err := http.Get(os.Getenv("API_ENDPOINT") + "/health")
	if get_err != nil {
		fmt.Println(get_err)
		return c.Render(http.StatusOK, "error.html", MakeData(nil))
	}

	body, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		log.Fatalln(read_err)
		return c.Render(http.StatusOK, "error.html", MakeData(nil))
	}

	var health objects.Health

	json.Unmarshal(body, &health)

	data := make(map[string]interface{})
	data["health"] = health

	return c.Render(http.StatusOK, "health.html", MakeData(data))
}

func Golink(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, os.Getenv("API_ENDPOINT")+"/"+c.Param("link"))
}

func Error(c echo.Context) error {
	return c.Render(http.StatusOK, "error.html", MakeData(nil))
}

func Endpoints(e *echo.Echo) {
	e.Static("/static", "static")
	e.GET("/", Hello)
	e.POST("/golink", CreateGolink)
	e.GET("/golinks", ListGolinks)
	e.GET("/health", HealthCheck)
	e.GET("/:link", Golink)
	e.GET("/error", Error)
}
