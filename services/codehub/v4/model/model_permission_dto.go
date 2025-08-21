package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PermissionDto struct {

	// **参数解释：** 排序id。
	Order *int32 `json:"order,omitempty"`

	// **参数解释：** 角色id。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleId *string `json:"role_id,omitempty"`

	// **参数解释：** 角色名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleName *string `json:"role_name,omitempty"`

	// **参数解释：** 角色中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleNameCn *string `json:"role_name_cn,omitempty"`

	// **参数解释：** 资源权限对象。
	ResourcePermissions map[string]ResourcePermissionDto `json:"resource_permissions,omitempty"`
}

func (o PermissionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionDto struct{}"
	}

	return strings.Join([]string{"PermissionDto", string(data)}, " ")
}
