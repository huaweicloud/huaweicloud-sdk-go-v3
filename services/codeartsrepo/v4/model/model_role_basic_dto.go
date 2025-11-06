package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RoleBasicDto struct {

	// **参数解释：** 角色唯一标识。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 角色名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 角色id。 **取值范围：** 字符串长度不少于1，不超过1000。
	RelatedRoleId *string `json:"related_role_id,omitempty"`

	// **参数解释：** 角色中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	ChineseName *string `json:"chinese_name,omitempty"`
}

func (o RoleBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleBasicDto struct{}"
	}

	return strings.Join([]string{"RoleBasicDto", string(data)}, " ")
}
