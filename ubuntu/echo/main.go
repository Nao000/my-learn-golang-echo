package main

import (
	"github.com/labstack/echo/v4"
	"echo/controller"
	"echo/core"
)

func main() {

	e := echo.New()

	// template を使う
	e.Renderer = core.GetTemplate()

	e.GET("/", controller.Index)
	e.GET("/hello", controller.Hello)

	e.Logger.Fatal(e.Start(":80"))
}