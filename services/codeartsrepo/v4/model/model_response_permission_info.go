package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponsePermissionInfo struct {

	// **参数解释：** 角色id **取值范围：** 不涉及。
	RoleId *string `json:"role_id,omitempty"`

	// **参数解释：** 角色名称 **取值范围：** 不涉及。
	RoleName *string `json:"role_name,omitempty"`

	// **参数解释：** 角色中文名称 **取值范围：** 不涉及。
	RoleNameCn *string `json:"role_name_cn,omitempty"`

	ResourcePermissions map[string]interface{} `json:"resource_permissions,omitempty"`
}

func (o ResponsePermissionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponsePermissionInfo struct{}"
	}

	return strings.Join([]string{"ResponsePermissionInfo", string(data)}, " ")
}
