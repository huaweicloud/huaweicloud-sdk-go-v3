package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InvalidVpnUser struct {

	// 用户名
	Name *string `json:"name,omitempty"`

	// 用户描述，0-64字符，中文、英文、数字包含下划线
	Description *string `json:"description,omitempty"`

	// 所属用户组名称
	UserGroupName *string `json:"user_group_name,omitempty"`

	// 失败原因
	Cause *string `json:"cause,omitempty"`
}

func (o InvalidVpnUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvalidVpnUser struct{}"
	}

	return strings.Join([]string{"InvalidVpnUser", string(data)}, " ")
}
