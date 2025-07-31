package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetAccountPasswordRequestBody 重置密码所需要的信息
type ResetAccountPasswordRequestBody struct {

	// 云服务厂商
	Vendor string `json:"vendor"`

	// 云服务
	ResourceProvider string `json:"resource_provider"`

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 实例类型
	InstanceType string `json:"instance_type"`

	// 需要改密的账号的状态，枚举值  TO_BE_CHANGED：待改密 FAILED：改密失败 SUCCESSFUL：改密成功 PROCESSING：改密中  不传默认修改所有状态的账号
	AccountsPasswordChangeStatus *[]string `json:"accounts_password_change_status,omitempty"`

	// 实例id列表
	ResourcesId []string `json:"resources_id"`
}

func (o ResetAccountPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetAccountPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ResetAccountPasswordRequestBody", string(data)}, " ")
}
