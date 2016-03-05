package main

//go:generate esc -o static.go -prefix public -ignore \.map$ public

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// addr is the listen address
var addr string

// debug enables debug mode, which uses local files
// instead of bundled ones
var debug bool

func init() {
	flag.StringVar(&addr, "addr", ":8080", "Address binding")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
}

func main() {
	flag.Parse()

	e := echo.New()

	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static content handler
	assetHandler := http.FileServer(FS(debug))

	// Web UI
	//e.Index("public/index.html")
	//e.Static("/public", "public")
	//e.ServeFile("/app/bundle.js", "public/bundle.js")
	e.Get("/", func(c *echo.Context) error {
		assetHandler.ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})

	// Serve webapp assets under the /app directory
	e.Get("/app/*", func(c *echo.Context) error {
		http.StripPrefix("/app", assetHandler).
			ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})

	// Listen to OS kill signals
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		fmt.Println("Exiting on signal")
		os.Exit(100)
	}()

	e.Run(addr)
}
