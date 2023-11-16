package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentInstallMode 组件详情
type ComponentInstallMode struct {

	// 组件名称
	Component string `json:"component"`

	// 该组件的角色部署信息
	NodeGroups []AssignedNodeGroup `json:"node_groups"`

	// 配置组件用户密码，该密码用于ClickHouse组件机机用户连接使用。 - 密码长度应在8～26个字符之间 - 不能与用户名或者倒序用户名相同 - 必须包含如下4种字符的组合 - 至少一个小写字母 - 至少一个大写字母 - 至少一个数字 - 至少一个特殊字符：!@$%^-_=+[{}]:,./?
	ComponentUserPassword *string `json:"component_user_password,omitempty"`

	// 配置组件default用户密码，该密码用于ClickHouse组件人机用户连接使用。 - 密码长度应在8～26个字符之间 - 不能与用户名或者倒序用户名相同 - 必须包含如下4种字符的组合 - 至少一个小写字母 - 至少一个大写字母 - 至少一个数字 - 至少一个特殊字符：!@$%^-_=+[{}]:,./?
	ComponentDefaultPassword *string `json:"component_default_password,omitempty"`
}

func (o ComponentInstallMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentInstallMode struct{}"
	}

	return strings.Join([]string{"ComponentInstallMode", string(data)}, " ")
}
