package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateInstanceConnectionRequestBody 用户登录连接测试信息。
type ValidateInstanceConnectionRequestBody struct {

	// 连接目标实例id
	TargetInstanceId *string `json:"target_instance_id,omitempty"`

	UserInfo *ValidateConnectionUserInfo `json:"user_info,omitempty"`
}

func (o ValidateInstanceConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateInstanceConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"ValidateInstanceConnectionRequestBody", string(data)}, " ")
}
