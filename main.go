package main

import (
	"gin-web/dddProject"
	"gin-web/dddProject/Infra/config"
)

func main() {
	// 加载server配置
	config.InitConfig("./config.ini")

	ser := dddProject.Server{}
	// 初始化启动
	ser.InitializedSystem("config/server_gin.xml")
	// 通过回调函数注册路由, 并启动服务
	ser.Run()
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
