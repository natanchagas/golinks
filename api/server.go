package main

import (
	"github.com/natanchagas/golinks/router"
)

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":1323"))
}
