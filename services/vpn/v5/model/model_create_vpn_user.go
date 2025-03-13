package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnUser struct {

	// 用户名
	Name *string `json:"name,omitempty"`

	// 用户密码
	Password *string `json:"password,omitempty"`

	// 用户描述，0-64字符，中文、英文、数字包含下划线
	Description *string `json:"description,omitempty"`

	// 所属用户组ID
	UserGroupName *string `json:"user_group_name,omitempty"`
}

func (o CreateVpnUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUser struct{}"
	}

	return strings.Join([]string{"CreateVpnUser", string(data)}, " ")
}
