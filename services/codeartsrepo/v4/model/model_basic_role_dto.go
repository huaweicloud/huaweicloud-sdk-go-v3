package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicRoleDto struct {

	// **参数解释：** 角色ID。 **取值范围：** 1-2147483647
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 角色名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  关联角色ID。 **取值范围：** 不涉及。
	RelatedRoleId *string `json:"related_role_id,omitempty"`

	// **参数解释：** 角色中文名。 **取值范围：** 不涉及。
	ChineseName *string `json:"chinese_name,omitempty"`
}

func (o BasicRoleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicRoleDto struct{}"
	}

	return strings.Join([]string{"BasicRoleDto", string(data)}, " ")
}
