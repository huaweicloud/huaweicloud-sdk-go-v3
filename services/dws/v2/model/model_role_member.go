package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RoleMember struct {

	// **参数解释**： 角色名。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 角色描述。 **取值范围**： 不涉及。
	Desc *string `json:"desc,omitempty"`

	// **参数解释**： 是否允许授予某个角色特定的权限。 **取值范围**： 不涉及。
	Permission *bool `json:"permission,omitempty"`

	// **参数解释**： 是否允许该角色将已获得的权限再转授给其他角色。 **取值范围**： 不涉及。
	GrantWith *bool `json:"grant_with,omitempty"`
}

func (o RoleMember) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleMember struct{}"
	}

	return strings.Join([]string{"RoleMember", string(data)}, " ")
}
