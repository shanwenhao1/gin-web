package redis

import (
	"encoding/json"
	"gin-web/server"
	"time"
)

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
	errR := server.GetRDS().SetNX(key, dataJson, 0).Err()
	return errR
}

func Set(key string, value interface{}) error {
	dataS, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var dataJson = string(dataS)
	errR := server.GetRDS().Set(key, dataJson, 0).Err()
	return errR
}

func SetWithTime(key string, value interface{}, timeOut time.Duration) error {
	dataS, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var dataJson = string(dataS)
	errR := server.GetRDS().Set(key, dataJson, timeOut).Err()
	return errR
}

func Get(key string, value interface{}) error {
	val, errR := server.GetRDS().Get(key).Result()
	if errR != nil {
		return errR
	}
	err := json.Unmarshal([]byte(val), value)
	return err
}

func Del(key string) error {
	errR := server.GetRDS().Del(key).Err()
	return errR
}

func IsNull(err error) bool {
	if err.Error() == "redis: nil" {
		return true
	}
	return false
}

func GetString(key string) (string, error) {
	return server.GetRDS().Get(key).Result()
}
