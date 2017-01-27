package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/utahta/go-webapp-proto/app/lib/dummy"
	"github.com/utahta/go-webapp-proto/app/lib/session"
	"github.com/utahta/go-webapp-proto/app/model"
)

func DummyIndex(c echo.Context) error {
	u, err := dummy.New().Do(3)
	if err != nil {
		return err
	}

	s := session.MustStart(c)
	hoge, ok := s.Get("hoge")
	if ok {
		s.Set("hoge", hoge.(int)+1)
	} else {
		s.Set("hoge", 1)
	}
	s.Save()

	v := struct{ User *model.User }{User: u}
	return c.Render(http.StatusOK, "assets/view/dummy/index.html", v)
}

func DummySearch(c echo.Context) error {
	s := session.MustStart(c)
	v := s.Session.Values["hoge"]

	log.Println(c.Path(), v)
	return c.Render(http.StatusOK, "assets/view/dummy/search.html", nil)
}
