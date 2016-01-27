package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index(c *echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func Eval(c *echo.Context) error {
	return c.Render(http.StatusOK, "eval.html", nil)
}

func PostEval(c *echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
