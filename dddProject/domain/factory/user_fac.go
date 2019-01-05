package factory

import "gin-web/dddProject/domain/model"

// factory创建User实例
func GetUser(userId string) model.User {
	// 如果角色不存在则创建(需持久化至数据库中), 否则查询User信息. 最后返回User实例
	return model.User{}
}
