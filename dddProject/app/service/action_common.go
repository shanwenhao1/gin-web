package service

import (
	"encoding/json"
	"gin-web/dddProject/Infra/enum"
	"gin-web/dddProject/Infra/log"
	"gin-web/dddProject/domain/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//响应数据模型
type ResponseJsonModel struct {
	Obj       interface{} `json:"obj"`       // 内容
	ErrorCode int         `json:"errorCode"` // 编码
	Token     interface{} `json:"token"`     // token
	ErrorMsg  interface{} `json:"errorMsg"`  // 消息
}

func GetRequestData(c *gin.Context, rjm interface{}) interface{} {
	var reqData model.RequestJsonModel
	req := c.Request
	addr := req.Header.Get("X-Real-IP") // 获取真实发出请求的客户端IP
	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For") // 获取IP(包含代理IP）
		if addr == "" {
			addr = req.RemoteAddr
		}
	}
	log.LogWithTag(log.InfoLog, log.ReqParse, "Request %s for %s", req.URL.Path, addr)
	dataS, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.LogWithTag(log.ErrorLog, log.ReqParse, "%w : %w", "Gin Read Body Error", err)
	}
	log.LogWithTag(log.InfoLog, log.ReqParse, "%v : %v", "The Request Body is", string(dataS))
	err = json.Unmarshal(dataS, rjm)
	if err != nil {
		log.LogWithTag(log.ErrorLog, log.ReqParse, "%v : %v", "Convert Body To Json Failed", err)
	}
	json.Unmarshal(dataS, &reqData)
	// you can do something with request obj
	if err != nil {
		ResponseData(c, GetDefaultRJM())
		return nil
	} else {
		return rjm
	}
}

/*
	响应函数
*/
func ResponseData(c *gin.Context, dataModel ResponseJsonModel) {
	c.JSON(http.StatusOK, dataModel)
}

//获取默认返回消息模型
func GetDefaultRJM(code ...int) ResponseJsonModel {
	if len(code) > 0 {
		return ResponseJsonModel{ErrorCode: code[0], ErrorMsg: enum.CodeMap[code[0]]}
	} else {
		return ResponseJsonModel{ErrorCode: enum.OPERATE_FAILED, ErrorMsg: enum.CodeMap[enum.OPERATE_FAILED]}
	}
}

//获取成功返回消息模型
func GetSuccessRJM(params ...interface{}) ResponseJsonModel {
	if len(params) == 1 {
		return ResponseJsonModel{ErrorCode: enum.OPERATE_SUCCESS, ErrorMsg: enum.CodeMap[enum.OPERATE_SUCCESS], Obj: params[0]}
	}
	if len(params) == 2 {
		return ResponseJsonModel{ErrorCode: enum.OPERATE_SUCCESS, ErrorMsg: enum.CodeMap[enum.OPERATE_SUCCESS], Obj: params[0], Token: params[1]}
	}
	return ResponseJsonModel{ErrorCode: enum.OPERATE_SUCCESS, ErrorMsg: enum.CodeMap[enum.OPERATE_SUCCESS]}
}

// 通用返回处理函数
func CommonResponse(c *gin.Context, model model.ParamModel) {
	if model.ErrorCode == 0 {
		ResponseData(c, GetSuccessRJM(model.Obj))
	} else {
		ResponseData(c, GetDefaultRJM(model.ErrorCode))
	}
}
