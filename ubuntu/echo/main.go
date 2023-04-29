package main

import (
	"github.com/labstack/echo/v4"
	"echo/core"
	"echo/router"
)

func main() {

	e := echo.New()

	// template を使う
	e.Renderer = core.GetTemplate()

	// ルーティング
	router.SetRouter(e)

	e.Logger.Fatal(e.Start(":80"))
}