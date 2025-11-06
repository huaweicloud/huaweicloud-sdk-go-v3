package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RoleSyncDto struct {

	// **参数解释：** 角色记录id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 角色id。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleId *string `json:"role_id,omitempty"`

	// **参数解释：** 角色同步开关。
	RoleSyncEnabled *bool `json:"role_sync_enabled,omitempty"`

	// **参数解释：** 角色名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleName *string `json:"role_name,omitempty"`

	// **参数解释：** 角色类型。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleType *string `json:"role_type,omitempty"`

	// **参数解释：** 角色中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleChineseName *string `json:"role_chinese_name,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o RoleSyncDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleSyncDto struct{}"
	}

	return strings.Join([]string{"RoleSyncDto", string(data)}, " ")
}
