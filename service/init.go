package service

import (
	"errors"
	"fmt"
	"log"

	"../model"

	_ "github.com/go-sql-driver/mysql" //下划线，仅执行init函数，不能调用包中的其他函数
	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

//var test string = "1111"

//
func init() {
	drivename := "mysql"
	dsName := "root:root@(127.0.0.1:3306)/im?charset=utf8"

	err := errors.New("")
	DbEngine, err = xorm.NewEngine(drivename, dsName)

	if nil != err && err.Error() != "" {
		log.Fatal(err.Error())
	}

	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(10)

	//自动建表
	DbEngine.Sync2(new(model.User), new(model.Contact), new(model.Community))

	fmt.Println("init db ok", DbEngine)

}
