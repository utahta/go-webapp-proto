package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type TemplateRenderer struct{}

func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// go-bindata を使って固めた html や js ファイルを取り出す
	// ただし、現状 Go の得意分野ではないので、軽めに使う想定
	src := append(
		MustAsset("view/base.html"),
		MustAsset(name)...,
	)
	tm := template.Must(template.New("base").Parse(string(src)))
	return tm.Execute(w, data)
}
