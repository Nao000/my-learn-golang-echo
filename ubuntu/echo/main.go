package main

import (
	"github.com/labstack/echo/v4"
	"echo/controller"

	"io" // template で必要
	"net/http" // template で必要
	"text/template" // template で必要
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", controller.Index)
	e.GET("/hello", Hello)
	e.Logger.Fatal(e.Start(":80"))
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "WorldAAAAAAAAAA")
}
