package router

import (
	"gin-web/dddProject/app/action"
	"github.com/gin-gonic/gin"
)

// 路由映射, 为用户请求服务入口, 映射至application层中服务
func Router(handleMap map[string]gin.HandlerFunc) {
	userAction := new(action.UserJsonModel)

	// 添加示例login路由
	handleMap["login"] = userAction.Login
}
