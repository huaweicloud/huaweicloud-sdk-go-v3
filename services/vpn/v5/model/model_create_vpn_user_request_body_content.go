package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVpnUserRequestBodyContent struct {

	// 用户名
	Name string `json:"name"`

	// 用户密码
	Password string `json:"password"`

	// 用户描述，0-64字符，中文、英文、数字包含下划线
	Description *string `json:"description,omitempty"`

	// 所属用户组ID
	UserGroupId *string `json:"user_group_id,omitempty"`
}

func (o CreateVpnUserRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateVpnUserRequestBodyContent", string(data)}, " ")
}
