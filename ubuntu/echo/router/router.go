package router

import (
	"github.com/labstack/echo/v4"
	"echo/controller"
)

func SetRouter(e *echo.Echo) {
	e.Static("/assets", "public/assets/")

	e.GET("/", controller.Index)
	e.GET("/hello", controller.Hello)
}