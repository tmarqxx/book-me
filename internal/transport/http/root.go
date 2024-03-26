package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRootServer() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	return e
}
