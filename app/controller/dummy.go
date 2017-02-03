package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/utahta/echo-sessions"
	"github.com/utahta/go-webapp-proto/app/lib/dummy"
	"github.com/utahta/go-webapp-proto/app/model"
)

func DummyIndex(c echo.Context) error {
	u, err := dummy.New().Do(3)
	if err != nil {
		return err
	}

	var hoge int
	if ok := sessions.MustGet(c, "hoge", &hoge); !ok {
		sessions.Set(c, "hoge", 1)
	} else {
		sessions.Set(c, "hoge", hoge+1)
	}
	sessions.Save(c)

	v := struct{ User *model.User }{User: u}
	return c.Render(http.StatusOK, "assets/view/dummy/index.html", v)
}

func DummySearch(c echo.Context) error {
	v := sessions.GetRaw(c, "hoge")

	log.Infof("%s %v", c.Path(), v)
	return c.Render(http.StatusOK, "assets/view/dummy/search.html", nil)
}
