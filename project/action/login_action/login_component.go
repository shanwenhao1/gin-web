package login_action

import (
	"gin-web/project/global"
	"github.com/gin-gonic/gin"
)

func (this Login) Login1(c *gin.Context) global.ParamModel {
	// do something
	return global.ParamModel{ErrorCode: global.OPERATE_SUCCESS}
}
