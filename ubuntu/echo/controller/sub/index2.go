package sub

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello2(c echo.Context) error {

	v_map := map[string]string{
		"example1_key": "example1_value",
		"example2_key": "example2_value",
		"example3_key": "example3_value",
		"example4_key": "example4_value",
	}

	return c.Render(http.StatusOK, "v", v_map)
}
