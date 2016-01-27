package renderer

import (
	"html/template"
	"io"
	"path"

	"github.com/labstack/echo"
)

type (
	Template struct {
		tmpls *template.Template
	}
)

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.tmpls.ExecuteTemplate(w, name, data)
}

func templates(base string, templates []string) *Template {
	tmpls := make([]string, len(templates))
	for i, tmpl := range templates {
		tmpls[i] = path.Join(base, tmpl+".html")
	}
	return &Template{
		tmpls: template.Must(template.ParseFiles(tmpls...)),
	}
}

func Set(e *echo.Echo, routes []string) *echo.Echo {
	e.SetRenderer(templates("view", routes))

	return e
}
