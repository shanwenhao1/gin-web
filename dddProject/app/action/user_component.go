package action

import (
	"fmt"
	"gin-web/dddProject/Infra/enum"
	"gin-web/dddProject/domain/model"
	"gin-web/dddProject/domain/service"
)

// 登录操作
func LoginH(req model.RequestJsonModel) model.ParamModel {
	// do something(不关乎领域逻辑和业务的)
	fmt.Println("-------")
	// 调用domain 内的User领域服务
	service.UserLogin(req)
	return model.ParamModel{ErrorCode: enum.OPERATE_SUCCESS}
}
