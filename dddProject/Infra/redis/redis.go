package redis

import (
	"encoding/json"
	"time"
	"io/ioutil"
	"encoding/xml"
	"github.com/go-redis/redis"
	"gin-web/dddProject/Infra/log"
)
var rds	*redis.Client

type RedisConfig struct {
	Redis_addr        string        `xml:"redis_addr"`
	Redis_passwd      string        `xml:"redis_passwd"`
	Redis_dbnum       int           `xml:"redis_dbnum"`
	Redis_Network     string        `xml:"redis_Network"`
	Redis_PoolSize    int           `xml:"redis_PoolSize"`
	Redis_IdleTimeout time.Duration `xml:"redis_IdleTimeout"`
}


func GetCacheKey(tag string, key string) string {
	return tag + "_" + key
}

// setNx如果key存在则不做任何操作, 无则插入
func SetNx(key string, value interface{}) error {
	dataS, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var dataJson = string(dataS)
	errR := rds.SetNX(key, dataJson, 0).Err()
	return errR
}

func Set(key string, value interface{}) error {
	dataS, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var dataJson = string(dataS)
	errR := rds.Set(key, dataJson, 0).Err()
	return errR
}

func SetWithTime(key string, value interface{}, timeOut time.Duration) error {
	dataS, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var dataJson = string(dataS)
	errR := rds.Set(key, dataJson, timeOut).Err()
	return errR
}

func Get(key string, value interface{}) error {
	val, errR := rds.Get(key).Result()
	if errR != nil {
		return errR
	}
	err := json.Unmarshal([]byte(val), value)
	return err
}

func Del(key string) error {
	errR := rds.Del(key).Err()
	return errR
}

func IsNull(err error) bool {
	if err.Error() == "redis: nil" {
		return true
	}
	return false
}

func GetString(key string) (string, error) {
	return rds.Get(key).Result()
}


func GetRDS() *redis.Client {
	return rds
}

// 初始化Redis数据源
func init() {
	dataS, rErr := ioutil.ReadFile("config/redisConfig.xml")
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
