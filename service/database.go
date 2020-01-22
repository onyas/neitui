package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/onyas/geekNews/model"
	"log"
	"os"
)

var DbEngine *xorm.Engine

func init() {
	DsName := os.Getenv("DATABASE_URL")
	if DsName == "" {
		log.Fatal("$DSNAME must be set")
	}

	driveName := "postgres"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driveName, DsName)
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
