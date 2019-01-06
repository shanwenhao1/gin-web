package db

import (
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"encoding/xml"
	"fmt"
	"gin-web/dddProject/domain/model"
	"gin-web/dddProject/Infra/log"
)

var (
	ds          gorm.DB
	tds         gorm.DB
)

type xmlStruct struct {
	DbName     string `xml:"dbname"`
	DbUser     string `xml:"dbuser"`
	DbUPwd     string `xml:"dbupwd"`
	DbUrl      string `xml:"dburl"`
	DbMaxConn  int    `xml:"dbmaxconn"`
	DbMaxIdle  int    `xml:"dbmaxidle"`
	DbLogModel bool   `xml:"dblogmodel"`
}

// 生产数据库
func GetDS() gorm.DB {
	return ds
}

// 测试数据库
func GetTDS() gorm.DB {
	return tds
}

//数据库连接初始化
func init() {
	dataS, rErr := ioutil.ReadFile("config/dbConfig.xml")
	if rErr != nil {
		log.LogWithTag(log.ErrorLog, log.InitSer, "读取数据库配置文件异常:[%v]", rErr)
		panic(rErr.Error())
	}
	configData := xmlStruct{}
	xErr := xml.Unmarshal(dataS, &configData)
	if xErr != nil {
		log.LogWithTag(log.ErrorLog, log.InitSer, "解析数据库配置文件异常:[%v]", xErr)
		panic(xErr.Error())
	}
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@%v/%v?charset=utf8&parseTime=True", configData.DbUser, configData.DbUPwd, configData.DbUrl, configData.DbName))
	if err != nil {
		panic(err.Error())
		log.LogWithTag(log.ErrorLog, log.InitSer, "初始化数据源异常:%v", err)
	} else {
		db.LogMode(configData.DbLogModel)
		db.SingularTable(true)
		db.DB().SetMaxOpenConns(configData.DbMaxConn)
		db.DB().SetMaxIdleConns(configData.DbMaxIdle)
		db.AutoMigrate(model.User{})
		ds = *db
		log.LogWithTag(log.InfoLog, log.InitSer, "数据源已初始化完成[最大打开连接数:%v,最大空闲连接数:%v]", configData.DbMaxConn, configData.DbMaxIdle)
	}
}