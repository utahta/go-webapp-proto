package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/utahta/go-webapp-proto/app/lib/config"
)

func NewEngine() (*xorm.Engine, error) {
	// プール的なところをさらっと書いてあげる必要がありそう
	// 都度接続でもいいかもしれないけど。
	src := fmt.Sprintf("%s@/%s?charset=utf8", config.C.Database.User, config.C.Database.DB)
	e, err := xorm.NewEngine("mysql", src)
	if err != nil {
		return nil, err
	}
	return e, nil
}
