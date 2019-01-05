package repository

import "gin-web/dddProject/domain/model"

// Repository(主要提供对象的访问和存储工作) 1: 面向对象资源库, 2: 面向持久化资源库

// 可将领域中user action逻辑中的一些查询或者存储的操作放入至此, 使得客户始终聚焦于模型
func ChangeName(user *model.User) {
	// 只是示例, 将
}
