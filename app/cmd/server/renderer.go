package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type TemplateRenderer struct{}

func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// go-bindata を使って固めた html や js ファイルを取り出す
	src, err := Asset(name)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	tm, err := template.New("webapp").Parse(string(src))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return tm.Execute(w, data)
}
