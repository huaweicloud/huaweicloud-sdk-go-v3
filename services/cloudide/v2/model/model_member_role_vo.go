package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberRoleVo struct {

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 插件数量
	ExtensionCount *int32 `json:"extension_count,omitempty"`

	// 角色名称
	Role *string `json:"role,omitempty"`

	// 角色ID
	RoleId *int64 `json:"role_id,omitempty"`

	// 发布商或插件ID
	RoleValue *string `json:"role_value,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`
}

func (o MemberRoleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberRoleVo struct{}"
	}

	return strings.Join([]string{"MemberRoleVo", string(data)}, " ")
}
