package main

import (
	"flag"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// mode is the execution mode
var mode = "dev"

var addr = ":9000"

func init() {
	flag.StringVar(&mode, "mode", "dev", "operation mode: 'dev' or 'prod'")
	flag.StringVar(&addr, "addr", ":9000", "Address binding")
}

func main() {
	flag.Parse()

	e := echo.New()

	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Web UI
	e.Index("public/index.html")
	e.Static("/public", "public")
	e.ServeFile("/app/bundle.js", "public/bundle.js")

	e.Run(addr)
}
