package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectAuthority **参数解释**： 对象权限信息。 **取值范围**： 不涉及。
type ObjectAuthority struct {

	// **参数解释**： 对象名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 角色权限集合。 **取值范围**： 不涉及。
	RoleAuthority *[]RoleAuthority `json:"role_authority,omitempty"`
}

func (o ObjectAuthority) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectAuthority struct{}"
	}

	return strings.Join([]string{"ObjectAuthority", string(data)}, " ")
}
