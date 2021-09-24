package config

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"server/model"
)

func GetConn() *gorm.DB {
	driverName := "mysql"
	DsName := "example:example@tcp(db:3306)/dothecasting?charset=utf8"
	err := errors.New("")
	db, err := gorm.Open(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	fmt.Println("init data base ok")
	db.LogMode(true)
	return db
}

func init() {
	db := GetConn()
	db.AutoMigrate(
		&model.Test{},
	)
}
