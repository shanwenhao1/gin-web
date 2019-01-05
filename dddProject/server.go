package dddProject

import (
	"encoding/xml"
	"fmt"
	"gin-web/dddProject/Infra/log"
	"gin-web/dddProject/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
)

type Server struct{}

type xmlStruct struct {
	DbName     string `xml:"dbname"`
	DbUser     string `xml:"dbuser"`
	DbUPwd     string `xml:"dbupwd"`
	DbUrl      string `xml:"dburl"`
	DbMaxConn  int    `xml:"dbmaxconn"`
	DbMaxIdle  int    `xml:"dbmaxidle"`
	DbLogModel bool   `xml:"dblogmodel"`
}

type RConfig struct {
	MPrefix     string `xml:"prefix"`
	MPort       string `xml:"port"`
	MEnv        string `xml:"serverModel"`
	MUploadPath string `xml:"uploadPath"`
	MSsl        bool   `xml:"useSSL"`
}

type RedisConfig struct {
	Redis_addr        string        `xml:"redis_addr"`
	Redis_passwd      string        `xml:"redis_passwd"`
	Redis_dbnum       int           `xml:"redis_dbnum"`
	Redis_Network     string        `xml:"redis_Network"`
	Redis_PoolSize    int           `xml:"redis_PoolSize"`
	Redis_IdleTimeout time.Duration `xml:"redis_IdleTimeout"`
}

var (
	ds          gorm.DB
	tds         gorm.DB
	rds         *redis.Client
	ConfigDataS RConfig
)

// 运行环境初始化(设置CPU核心数、读取server端口配置等)
func (this Server) InitializedSystem(path string) {
	dataS, rErr := ioutil.ReadFile(path)
	if rErr != nil {
		log.LogWithTag(log.ErrorLog, log.InitSer, "读取服务配置文件异常:[%v]", rErr)
		panic(rErr.Error())
	}
	configData := RConfig{}
	xErr := xml.Unmarshal(dataS, &configData)
	if xErr != nil {
		log.LogWithTag(log.ErrorLog, log.InitSer, "解析服务库配置文件异常:[%v]", xErr)
		panic(xErr.Error())
	}
	ConfigDataS = configData
	cc := runtime.NumCPU()
	// running in multi core
	runtime.GOMAXPROCS(cc)
	log.LogWithTag(log.InfoLog, log.InitSer, "运行环境初始化完成[处理器核心数:%d]", cc)
}

// 网络框架初始化(注册回调函数供加载路由)
func (this Server) InitializedNetwork(tls bool, callback func(handlerMap map[string]gin.HandlerFunc)) {
	handlerMap := make(map[string]gin.HandlerFunc)
	// 设置为线上环境
	gin.SetMode(gin.ReleaseMode)
	//router := gin.Default()
	router := gin.New()
	callback(handlerMap)
	log.LogWithTag(log.InfoLog, log.InitSer, "http网络框架初始化完成[%s][%s]", ConfigDataS.MPort, ConfigDataS.MEnv)
	for patten, handle := range handlerMap {
		// 对特殊的路由处理, 其余一律采用POST方法
		if strings.Contains(patten, "Upload") {
			log.LogWithTag(log.InfoLog, log.InitSer, "文件下载插件注册完成[%s]", patten)
			//router.Any("/test/Upload/", handle)
		} else {
			router.POST(ConfigDataS.MPrefix+patten, handle)
		}
	}
	// 另一种注册路由方式
	router.Static("/test/Upload/", "./Upload")
	// router.RunTLS使用HTTPS加密连接(需生成ssl key), router.Run使用http连接
	if ConfigDataS.MSsl {
		router.RunTLS(ConfigDataS.MPort, "keys/my.pem", "keys/my.key")
	} else {
		fmt.Println("xxxxxxxxxxx Server Run On", ConfigDataS.MPort, "...")
		router.Run(ConfigDataS.MPort)
	}
}

// 初始化Redis数据源
func (this Server) InitializedRedisSource(path string) {
	dataS, rErr := ioutil.ReadFile(path)
	if rErr != nil {
		log.LogWithTag(log.ErrorLog, log.InitSer, "读取Redis服务配置文件异常:[%v]", rErr)
		panic(rErr.Error())
	}
	reConfig := RedisConfig{}
	xErr := xml.Unmarshal(dataS, &reConfig)
	if xErr != nil {
		log.LogWithTag(log.ErrorLog, log.InitSer, "解析Redis配置文件异常:[%v]", xErr)
		panic(xErr.Error())
	}
	client := redis.NewClient(&redis.Options{
		Addr:        reConfig.Redis_addr,
		Password:    reConfig.Redis_passwd, // no password set
		DB:          reConfig.Redis_dbnum,  // use default DB
		Network:     reConfig.Redis_Network,
		PoolSize:    reConfig.Redis_PoolSize,
		IdleTimeout: reConfig.Redis_IdleTimeout,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	rds = client
	log.LogWithTag(log.InfoLog, log.InitSer, "Redis已初始化完成[%v|%v]", reConfig.Redis_addr, reConfig.Redis_dbnum)
}

//数据库框架初始化
func (this Server) InitializedDataSource(path string) {
	dataS, rErr := ioutil.ReadFile(path)
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

func GetDS() gorm.DB {
	return ds
}

func GetTDS() gorm.DB {
	return tds
}

func GetRDS() *redis.Client {
	return rds
}
