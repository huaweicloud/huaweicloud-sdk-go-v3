package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlterRoleInput 角色信息
type AlterRoleInput struct {

	// 描述信息。最大长度为4000个字符。
	Description *string `json:"description,omitempty"`

	// 角色名称。只能包含字母、数字和下划线，且长度为1~255个字符。
	RoleName string `json:"role_name"`
}

func (o AlterRoleInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlterRoleInput struct{}"
	}

	return strings.Join([]string{"AlterRoleInput", string(data)}, " ")
}
