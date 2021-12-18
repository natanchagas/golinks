package main

import (
	"github.com/natanchagas/golinks/front/router"
)

func main() {

	e := router.New()
	e.Logger.Fatal(e.Start(":80"))
}
