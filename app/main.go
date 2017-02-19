package main

import (
	"github.com/boj/redistore"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/utahta/echo-sessions"
	"github.com/utahta/go-webapp-proto/app/assets"
	"github.com/utahta/go-webapp-proto/app/controller"
	"github.com/utahta/go-webapp-proto/app/lib/config"
	"github.com/utahta/go-webapp-proto/app/lib/db"
)

// GOPATH 下に置いて開発
// GOPATH 下に移動して git clone

func run() error {
	// 設定ファイルを読み込み
	if err := config.Load("dev", "config/local"); err != nil {
		return err
	}

	if err := db.Open(); err != nil {
		return err
	}
	defer db.Close()

	// セッション用途
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		return err
	}
	defer store.Close()

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG) // ログレベルの設定

	// ミドルウェア設定
	// e.Pre(), e.Use() がある詳しくはドキュメント参照
	e.Use(middleware.Recover())                     // パニックが起きたとき、リカバーしてエラーレスポンスを返す
	e.Use(middleware.Logger())                      // リクエスト情報をログに書き出す。default stdout
	e.Use(sessions.Sessions("WEBAPPSESSID", store)) // セッションは自前ミドルウェア
	e.Renderer = new(assets.Template)               // テンプレートレンダラー

	// 静的ファイル
	e.GET("/public/*", assets.FileServerHandler())

	// ルーティング
	g := e.Group("/dummy")
	g.GET("/", controller.DummyIndex)
	g.GET("/search", controller.DummySearch)

	return e.Start(":8888")
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
