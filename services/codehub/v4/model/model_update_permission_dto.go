package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePermissionDto struct {

	// **参数解释：** 角色id **取值范围：** 不涉及。
	RoleId *string `json:"role_id,omitempty"`

	// **参数解释：** 角色名称 **取值范围：** 不涉及。
	RoleName *string `json:"role_name,omitempty"`

	// **参数解释：** 更新权限点详情 **取值范围：** 不涉及。
	Permissions *[]UpdatePermissionDetail `json:"permissions,omitempty"`
}

func (o UpdatePermissionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermissionDto struct{}"
	}

	return strings.Join([]string{"UpdatePermissionDto", string(data)}, " ")
}
