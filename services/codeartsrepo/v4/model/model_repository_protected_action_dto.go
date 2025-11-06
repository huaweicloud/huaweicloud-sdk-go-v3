package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryProtectedActionDto struct {

	// **参数解释：** 事件名称。 **取值范围：** 不涉及。
	Action *string `json:"action,omitempty"`

	// **参数解释：** 是否启用。 **取值范围：** - true，开启规则限制。 - false，未开启规则限制。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 白名单用户列表。 **取值范围：** 不涉及。
	Users *[]RepositoryUserDto `json:"users,omitempty"`

	// **参数解释：** 白名单用户组列表。 **取值范围：** 不涉及。
	UserTeams *[]UserTeamBasicDto `json:"user_teams,omitempty"`

	// **参数解释：** 白名单角色列表。 **取值范围：** 不涉及。
	Roles *[]BasicRoleDto `json:"roles,omitempty"`
}

func (o RepositoryProtectedActionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryProtectedActionDto struct{}"
	}

	return strings.Join([]string{"RepositoryProtectedActionDto", string(data)}, " ")
}
