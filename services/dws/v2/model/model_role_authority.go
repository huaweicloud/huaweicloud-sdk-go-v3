package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleAuthority **参数解释**： 角色权限信息。 **取值范围**： 不涉及。
type RoleAuthority struct {

	// **参数解释**： 角色名称。 **取值范围**： 不涉及。
	Role *string `json:"role,omitempty"`

	// **参数解释**： 权限列表。 **取值范围**： 不涉及。
	RightList *[]string `json:"right_list,omitempty"`
}

func (o RoleAuthority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleAuthority struct{}"
	}

	return strings.Join([]string{"RoleAuthority", string(data)}, " ")
}
