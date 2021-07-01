package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/natanchagas/golinks/objects"
)

var pgbd objects.Pgdb = objects.Pgdb{
	Host:     "localhost",
	User:     "golinksadm",
	Password: "mdasknilog",
	Dbname:   "golinks",
	Port:     5432,
	Sslmode:  "disable",
}

func HealthCheck(c echo.Context) error {

	var status int
	var health objects.Health

	pgconn, err := pgbd.Init()
	if err != nil {
		health.Database = "Down"
		health.Status = strconv.Itoa(http.StatusPartialContent)
		status = http.StatusPartialContent
	} else {
		health.Database = "Up"
		health.Status = strconv.Itoa(http.StatusOK)
		status = http.StatusOK
	}

	health.Api = "Up"

	pgconn.CloseConnection()
	return c.JSON(status, health)
}

func CreateGolink(c echo.Context) error {

	pgconn, err := pgbd.Init()
	if err != nil {
		panic(err)
	}

	var g objects.Golink
	c.Bind(&g)
	fmt.Println(g.OriginalUrl)

	query_result := pgconn.GetByOriginalUrl(g.OriginalUrl)
	fmt.Println(query_result)

	defer pgconn.CloseConnection()

	if query_result.Golink == "" {
		id := strings.Split(uuid.New().String(), "-")[0]
		query_result = pgconn.GetByGolink(query_result.Golink)
		if query_result.Golink == id {
			return c.JSON(http.StatusOK, query_result)
		} else {
			time := time.Now().Format(time.RFC3339)
			insert := objects.Golink{OriginalUrl: g.OriginalUrl, Golink: id, CreationDate: time}
			pgconn.CreateGolink(&insert)
			return c.JSON(http.StatusOK, insert)
		}
	} else {
		return c.JSON(http.StatusOK, query_result)
	}

}

func ListGolinks(c echo.Context) error {

	pgconn, err := pgbd.Init()
	if err != nil {
		panic(err)
	}

	golink := pgconn.GetAll()

	pgconn.CloseConnection()

	return c.JSON(http.StatusOK, &golink)
}

func Endpoints(e *echo.Echo) {
	e.GET("/health", HealthCheck)
	e.POST("/golink", CreateGolink)
	e.GET("/golinks", ListGolinks)
	// e.GET("/golink/:link", handlers.AddCat)
}
