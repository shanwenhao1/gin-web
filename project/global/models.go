package global

import "gin-web/project/models"

//异常消息模型
type ParamModel struct {
	ErrorCode int
	ErrorMsg  interface{}
	Obj       interface{}
	OtherData interface{}
}

//请求数据模型
type RequestJsonModel struct {
	AppId         string      `json:"appId"`
	Token         string      `json:"token"`
	Obj           interface{} `json:"obj"`
	ClientType    string      `json:"clientType"`
	Sign          string      `json:"sign"`
	TimeStamp     string      `json:"time_stamp"`
	ClientVersion string      `json:"clientVersion"`
}

// 请求测试模型1
type LoginJsonModel struct {
	RequestJsonModel
	Obj models.User `json:"obj"`
}
