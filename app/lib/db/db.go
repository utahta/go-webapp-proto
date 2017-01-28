package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/utahta/go-webapp-proto/app/lib/config"
)

/**
 * db.Open はサーバを立ち上げるときに呼ぶ
 *
 * db.E を使いまわしていく
 * e.g.
 *   db.E.Where(...).Get(...)
 */
var E *xorm.Engine

func Open() (err error) {
	src := fmt.Sprintf("%s@/%s?charset=utf8", config.C.Database.User, config.C.Database.DB)
	E, err = xorm.NewEngine(config.C.Database.Driver, src)
	if err != nil {
		return err
	}

	// 最大コネクション数などはこのあたりで設定

	return nil
}
