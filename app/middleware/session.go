package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/utahta/go-webapp-proto/app/lib/session"
)

// 詳細はゴリラを使うのでよしとして、多少自前で書いてあげる必要

func Session(name string, store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session.New(c, name, store)
			return next(c)
		}
	}
}
