package router

import (
	"gin-web/project/action/login_action"
	"github.com/gin-gonic/gin"
)

// 路由映射
func Router(handleMap map[string]gin.HandlerFunc) {
	loginAction := new(login_action.Login)

	// 添加示例login路由
	handleMap["login"] = loginAction.Login
}
