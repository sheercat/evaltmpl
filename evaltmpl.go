package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"path"

	"html/template"
	"net/http"

	// _ "github.com/golang/glog"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var portNumber = flag.Int("port", 3000, "port number.")

type (
	Template struct {
		tmpls *template.Template
	}
)

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.tmpls.ExecuteTemplate(w, name, data)
}

func indexPage(c *echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "Joe")
}

func evalPage(c *echo.Context) error {
	return c.Render(http.StatusOK, "eval.html", "Joe")
}

func postEvalPage(c *echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// Provide cached templates
func templates(base string, templates []string) *Template {
	tmpls := make([]string, len(templates))
	for i, tmpl := range templates {
		tmpls[i] = path.Join(base, tmpl+".html")
	}
	return &Template{
		tmpls: template.Must(template.ParseFiles(tmpls...)),
	}
}

func main() {
	flag.Parse()
	e := echo.New()
	e.Debug()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	e.Favicon("static/favicon.ico")
	e.Static("/", "static")

	// get cached tempalte
	e.SetRenderer(templates("view", []string{"index", "eval"}))

	e.Get("/", indexPage)
	e.Get("/eval/", evalPage)
	e.Post("/eval/", postEvalPage)

	log.Println(fmt.Sprintf("listen :%d", *portNumber))
	e.Run(fmt.Sprintf(":%d", *portNumber))
}
