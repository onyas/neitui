package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/onyas/geekNews/model"
	"log"
)

var DbEngine *xorm.Engine

func init() {
	drivename := "postgres"
	DsName := "postgres://postgres:root@127.0.0.1:5432/chat?sslmode=disable"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(drivename, DsName)
	if nil != err && "" != err.Error() {
		log.Println(err.Error())
	}
	//是否显示SQL语句
	DbEngine.ShowSQL(false)
	//数据库最大打开的连接数
	DbEngine.SetMaxOpenConns(2)

	//自动Sync
	err = DbEngine.Sync2(new(model.JobInfo))
	if nil != err {
		log.Println(err.Error())
	}
	fmt.Println("init data base ok")
}
