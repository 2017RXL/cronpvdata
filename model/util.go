package model

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var MySqlDB *gorm.DB

func init()  {
	initLog()
	var err error
	MySqlDB, err = gorm.Open("mysql", "wp:123456@tcp(172.17.118.245:3306)/wp?charset=utf8&parseTime=true")
	if err != nil {
		Logger.Info("connet database err=%s", err)
	}
	MySqlDB.SingularTable(true)
	Logger.Info("db %s", MySqlDB)
	MySqlDB.AutoMigrate(&GridPredInfo{})
	var gri GridPredInfo
	boo:=MySqlDB.HasTable(gri)
	Logger.Info("该模型是否存在 ",boo)
}

var Logger *logs.BeeLogger
func initLog() {
	Logger = logs.NewLogger()
	Logger.SetLogger(logs.AdapterMultiFile, `{"filename":"cronpvdata.log","maxdays":10,"color":true}`)
	Logger.EnableFuncCallDepth(true)
	Logger.SetLogFuncCallDepth(2)
	Logger.EnableFuncCallDepth(true)
}
