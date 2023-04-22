package controller

import(
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}

func Hello(c echo.Context) error {

	return c.Render(http.StatusOK, "sub/welcome", "World 2")
}