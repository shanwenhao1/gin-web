package action

import (
	"gin-web/dddProject/app/service"
	"gin-web/dddProject/domain/model"
	"github.com/gin-gonic/gin"
)

// user action 的请求json参数
type UserJsonModel struct {
	model.RequestJsonModel
	Obj model.User `json:"obj"`
}

// 登录请求
func (this UserJsonModel) Login(c *gin.Context) {
	// 获取请求参数, 可以考虑在此添加中间件
	rjm := service.GetRequestData(c, &model.RequestJsonModel{})
	if rjm != nil {
		jsonModel := *rjm.(*model.RequestJsonModel)
		// 可以做一些验证
		result := LoginH(jsonModel)
		service.CommonResponse(c, result)
	}
}
