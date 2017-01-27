package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/boj/redistore"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/utahta/go-webapp-proto/app/controller"
	"github.com/utahta/go-webapp-proto/app/lib/config"
	appmiddleware "github.com/utahta/go-webapp-proto/app/middleware"
)

// GOPATH の下に置いて開発想定
// 一発目は go get github.com/utahta/go-webapp-proto してもらう

type Template struct{}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// html や js などを go-bindata を使って assets.go にしている
	// そこからよしなに取り出す
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

func doMain() error {
	// 設定ファイルを読み込み
	// とりあえず直値をいれてる。CLI フラグで渡すのがいいかもしれない
	if err := config.Load("dev", "../../../config/local"); err != nil {
		return err
	}

	// セッション用の redis に繋ぐ
	store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		return err
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG) // ログレベルの設定
	e.Use(middleware.Recover()) // パニックになったときにリカバーしてレスポンスを返す
	e.Use(middleware.Logger())  // リクエスト情報をログ吐き出し。default stdout なのでデバッグ用かな。
	// 自前のミドルウェアを差し込むことも可能
	// e.Pre(), e.Use() がある詳しくはドキュメント参照
	e.Use(appmiddleware.Session("WEBAPPSESSID", store))

	// template をよしなに解釈
	e.Renderer = new(Template)

	// ルーティング
	e.GET("/static/*", controller.StaticPublic)

	// グループ化して云々もできそう
	g := e.Group("/dummy")
	g.GET("/", controller.DummyIndex)
	g.GET("/search", controller.DummySearch)

	// 管理画面 を app と同じ階層でやるか、admin を app の隣につくるか要検討

	return e.Start(":1323")
}

func main() {
	if err := doMain(); err != nil {
		log.Fatal(err)
	}
}
