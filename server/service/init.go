package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"server/model"
)

var DbEngine *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "example:example@tcp(db:3306)/dothecasting?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Test))
	fmt.Println("init data base ok")
}
