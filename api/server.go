package main

import (
	"github.com/natanchagas/golinks/router"
)

// @title Golinks
// @version 0.0.1
// @description Welcome to the Golinks API. Golinks is an URL shortener.\n\nThis is a project to learn and understand the creations of APIs with Go. The framework of choice was [Echo](https://echo.labstack.com/), for no special reason. You can take a look on the full source code at [Github](https://github.com/natanchagas/golinks).\n\nYou can check my other projects at [https://github.com/natanchagas](https://github.com/natanchagas).
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host golinks.natanchagas.com
// @BasePath /api

func main() {

	e := router.New()
	e.Logger.Fatal(e.Start(":80"))
}
