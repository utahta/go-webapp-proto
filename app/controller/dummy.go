package controller

import (
	"net/http"

	"github.com/utahta/go-webapp-proto/app/lib/dummy"
	"github.com/utahta/go-webapp-proto/app/lib/render"
	"github.com/utahta/go-webapp-proto/app/model"
)

func DummyIndex(w http.ResponseWriter, r *http.Request) {
	u, err := dummy.New().Do(3)
	if err != nil {
		render.InternalServerError(w, r)
		return
	}

	//var hoge int
	//if ok := sessions.MustGet(c, "hoge", &hoge); !ok {
	//	sessions.Set(c, "hoge", 1)
	//} else {
	//	sessions.Set(c, "hoge", hoge+1)
	//}
	//sessions.Save(c)

	v := struct{ User *model.User }{User: u}
	render.Template(w, r, "dummy/index", v)
}

func DummySearch(w http.ResponseWriter, r *http.Request) {
	//v := sessions.GetRaw(c, "hoge")

	render.Template(w, r, "dummy/search", nil)
}
