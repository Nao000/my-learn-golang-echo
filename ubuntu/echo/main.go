package main

import (
	"github.com/labstack/echo/v4"
	"echo/controller"
)

func main() {
	e := echo.New()
	e.GET("/", controller.Index)
	e.Logger.Fatal(e.Start(":80"))
}