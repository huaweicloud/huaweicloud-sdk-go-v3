package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckCanAuthUsersOfAppRequest struct {
	// 应用ID

	AppId string `json:"app_id"`
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 过滤条件 - 过滤条件，未提供时返回包括应用成员在内的所有候选用户列表 - 取值members时，过滤掉当前应用成员，适合增量添加应用成员时使用

	Filter *string `json:"filter,omitempty"`
}

func (o CheckCanAuthUsersOfAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCanAuthUsersOfAppRequest struct{}"
	}

	return strings.Join([]string{"CheckCanAuthUsersOfAppRequest", string(data)}, " ")
}
