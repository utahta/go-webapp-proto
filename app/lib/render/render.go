package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/pressly/chi/render"
	"github.com/utahta/go-webapp-proto/app/assets"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusInternalServerError)
	render.PlainText(w, r, "")
}

func PlainText(w http.ResponseWriter, r *http.Request, v string) {
	render.PlainText(w, r, v)
}

func JSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.JSON(w, r, v)
}

func Template(w http.ResponseWriter, r *http.Request, name string, v interface{}) {
	src := append(
		assets.MustAsset("view/base.html"),
		assets.MustAsset(fmt.Sprintf("view/%s.html", name))...,
	)

	buf := new(bytes.Buffer)
	tm := template.Must(template.New("base").Parse(string(src)))
	if err := tm.Execute(buf, v); err != nil {
		InternalServerError(w, r)
		return
	}

	render.Status(r, http.StatusOK)
	render.HTML(w, r, buf.String())
}
