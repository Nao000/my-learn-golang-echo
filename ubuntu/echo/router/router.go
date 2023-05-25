package router

import (
	"github.com/labstack/echo/v4"
	"echo/controller"
	controller_sub "echo/controller/sub"
	controller_accessdb "echo/controller/accessdb"
	controller_registerdb "echo/controller/registerdb"
	controller_transactiondb "echo/controller/transactiondb"
)

func SetRouter(e *echo.Echo) {
	e.Static("/assets", "public/assets/")

	e.GET("/", controller.Index)
	e.GET("/hello", controller.Hello)
	e.GET("/hello/sub", controller_sub.Hello)
	e.GET("/hello/sub2", controller_sub.Hello2)
	e.GET("/accessdb", controller_accessdb.Index)
	e.POST("/registerdb", controller_registerdb.Index)
	e.POST("/transactiondb", controller_transactiondb.Index)
}