package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/utahta/go-webapp-proto/app/lib/config"
)

var engine *xorm.Engine

func E() *xorm.Engine {
	return engine
}

func Open() (err error) {
	src := fmt.Sprintf("%s@/%s?charset=utf8", config.C.Database.User, config.C.Database.DB)
	engine, err = xorm.NewEngine(config.C.Database.Driver, src)
	if err != nil {
		return err
	}

	// db connection settings...

	return nil
}

func Close() {
	engine.Close()
}
