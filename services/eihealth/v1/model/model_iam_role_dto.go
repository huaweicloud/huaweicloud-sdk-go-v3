package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IamRoleDto struct {

	// 权限ID。
	Id *string `json:"id,omitempty"`

	// 权限显示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 权限类型：AX为全局级角色，XA为项目级角色。
	Type *string `json:"type,omitempty"`

	// 权限描述。
	Description *string `json:"description,omitempty"`

	// 授权项，指对资源的具体操作权限。
	Actions *[]string `json:"actions,omitempty"`
}

func (o IamRoleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamRoleDto struct{}"
	}

	return strings.Join([]string{"IamRoleDto", string(data)}, " ")
}
