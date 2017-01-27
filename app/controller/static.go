package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func StaticPublic(c echo.Context) error {
	return c.Render(http.StatusOK, fmt.Sprintf("assets/public/%s", c.Param("*")), nil)
}
