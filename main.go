package main

import (
	"gin-web/project/router"
	"gin-web/server"
	"gin-web/tool/config"
	"gin-web/tool/log"
)

func main() {
	// 加载server配置
	config.InitConfig("./config.ini")
	// 日志初始化
	log.InitializedLog4go("config/log4go.xml")

	ser := server.Server{}
	// 初始化启动
	ser.InitializedSystem("config/server_gin.xml")
	// 初始化数据库映射
	ser.InitializedDataSource("config/dbConfig.xml")
	// 初始化redis
	ser.InitializedRedisSource("config/redisConfig.xml")
	// 通过回调函数注册路由, 并启动服务
	ser.InitializedNetwork(false, router.Router)
	//router := gin.Default()
	//s := &http.Server{
	//	Addr:           ":8080",
	//	Handler:        router,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()
}
