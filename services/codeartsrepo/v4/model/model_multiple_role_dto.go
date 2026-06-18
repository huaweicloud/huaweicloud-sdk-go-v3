package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultipleRoleDto struct {

	// **参数解释：** 角色id。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleId *string `json:"role_id,omitempty"`

	// **参数解释：** 角色名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleName *string `json:"role_name,omitempty"`

	// **参数解释：** 角色中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleChineseName *string `json:"role_chinese_name,omitempty"`
}

func (o MultipleRoleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultipleRoleDto struct{}"
	}

	return strings.Join([]string{"MultipleRoleDto", string(data)}, " ")
}
