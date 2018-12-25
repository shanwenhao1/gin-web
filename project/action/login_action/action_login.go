package login_action

import (
	"gin-web/project/action"
	"gin-web/project/conponent"
	"gin-web/project/global"
	"github.com/gin-gonic/gin"
)

type Login struct {
	UserName string
	PassWord string
}

func (this Login) Login(c *gin.Context) {
	// 获取请求参数, 可以考虑在此添加中间件
	rjm := action.GetRequestData(c, &global.LoginJsonModel{})
	if rjm != nil {
		jsonModel := *rjm.(*global.RequestJsonModel)
		// 可以做一些验证
		model := new(conponent.LoginHandle).Login(jsonModel)
		action.CommonResponse(c, model)
	}
}
