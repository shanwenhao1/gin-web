package conponent

import (
	"gin-web/project/global"
)

type LoginHandle struct{}

// 登录实际处理函数
func (this LoginHandle) Login(rjm global.RequestJsonModel) global.ParamModel {
	// do something
	return global.ParamModel{ErrorCode: global.OPERATE_SUCCESS}
}
