package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccountPasswordRequestBody 回写改密结果所需要的信息
type UpdateAccountPasswordRequestBody struct {

	// 云服务厂商
	Vendor string `json:"vendor"`

	// 云服务
	ResourceProvider string `json:"resource_provider"`

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 实例类型
	InstanceType string `json:"instance_type"`

	// 改密结果
	PasswordChangeResult []UpdateAccountPassword `json:"password_change_result"`
}

func (o UpdateAccountPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccountPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAccountPasswordRequestBody", string(data)}, " ")
}
