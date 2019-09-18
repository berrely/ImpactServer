package middleware

import (
	"github.com/labstack/echo"
	"strconv"
)

func Cache(maxAge int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "max-age="+strconv.Itoa(maxAge))
			return next(c)
		}
	}
}