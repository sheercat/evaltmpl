package router

import (
	"github.com/labstack/echo"
	"github.com/sheercat/evaltmpl/controller"
)

var (
	Routes = []string{"index", "eval"}
)

func Set(e *echo.Echo) *echo.Echo {

	e.Favicon("static/favicon.ico")
	e.Static("/", "static")

	e.Get("/", controller.Index)
	e.Get("/eval/", controller.Eval)
	e.Post("/eval/", controller.PostEval)

	return e
}
