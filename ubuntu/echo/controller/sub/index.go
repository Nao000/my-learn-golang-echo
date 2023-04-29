package sub

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) error {

	return c.Render(http.StatusOK, "sub/welcome", "From controller/sub/index")
}