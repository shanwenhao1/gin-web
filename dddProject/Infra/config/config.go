package config

import (
	"github.com/go-ini/ini"
	"log"
	"strings"
)

var conf *ini.File

func InitConfig(config string) {

	var err error
	conf, err = ini.Load(config)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("加载配置文件成功")
}

func Get(str string) string {

	strArr := strings.Split(str, ".")

	if len(strArr) == 2 {
		return conf.Section(strArr[0]).Key(strArr[1]).String()
	}

	return conf.Section("").Key(strArr[0]).String()
}
