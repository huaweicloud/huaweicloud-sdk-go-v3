package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckAuthUsersOfAppRequest struct {
	// 应用ID

	AppId string `json:"app_id"`
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 查询应用的指定名称的成员，精确匹配

	UserName *string `json:"user_name,omitempty"`
}

func (o CheckAuthUsersOfAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAuthUsersOfAppRequest struct{}"
	}

	return strings.Join([]string{"CheckAuthUsersOfAppRequest", string(data)}, " ")
}
