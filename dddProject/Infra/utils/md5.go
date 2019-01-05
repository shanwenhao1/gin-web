package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const (
	key = "" // generate by uuid
)

// 创建会话时, 验证Md5
func SessionMd5(timeStamp int) (string, string) {
	timeStr := strconv.Itoa(timeStamp)
	seSs := key + timeStr
	data := []byte(seSs)
	hmd5 := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", hmd5)
	return md5Str, seSs
}

func VerifySeSs(md5 string, timeStamp int) bool {
	serMd5, _ := SessionMd5(timeStamp)
	if serMd5 != md5 {
		return false
	}
	return true
}
