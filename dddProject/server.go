package dddProject

import (
	"encoding/xml"
	"gin-web/dddProject/Infra/log"
	"io/ioutil"
	"runtime"
	"gin-web/dddProject/interfaces/router"
)

type Server struct{}



// 运行环境初始化(设置CPU核心数、读取server端口配置等)
func (this Server) InitializedSystem(path string) {
	dataS, rErr := ioutil.ReadFile(path)
	if rErr != nil {
		log.LogWithTag(log.ErrorLog, log.InitSer, "读取服务配置文件异常:[%v]", rErr)
		panic(rErr.Error())
	}
	configData := router.RConfig{}
	xErr := xml.Unmarshal(dataS, &configData)
	if xErr != nil {
		log.LogWithTag(log.ErrorLog, log.InitSer, "解析服务库配置文件异常:[%v]", xErr)
		panic(xErr.Error())
	}
	router.ConfigDataS = configData
	cc := runtime.NumCPU()
	// running in multi core
	runtime.GOMAXPROCS(cc)
	log.LogWithTag(log.InfoLog, log.InitSer, "运行环境初始化完成[处理器核心数:%d]", cc)
}

// 服务启动
func (this Server) Run(){
	router.Init()
}
