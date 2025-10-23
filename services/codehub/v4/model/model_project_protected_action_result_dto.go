package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectProtectedActionResultDto struct {

	// **参数解释：** 动作。 **取值范围：** 字符串长度不少于1，不超过1000。
	Action *string `json:"action,omitempty"`

	// **参数解释：** 是否开启。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 权限点。 **取值范围：** 字符串长度不少于1，不超过1000。
	Levels *string `json:"levels,omitempty"`

	// **参数解释：** 权限名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	LevelNames *string `json:"level_names,omitempty"`

	// **参数解释：** 成员列表。
	Users *[]UserBasicDto `json:"users,omitempty"`

	// **参数解释：** 成员组列表。
	UserTeams *[]UserTeamBasicDto `json:"user_teams,omitempty"`

	// **参数解释：** 角色列表。
	Roles *[]RoleBasicDto `json:"roles,omitempty"`

	// **参数解释：** 操作选择列表。
	AdditionSwitchers *[]ForceActionEnableDto `json:"addition_switchers,omitempty"`
}

func (o ProjectProtectedActionResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectProtectedActionResultDto struct{}"
	}

	return strings.Join([]string{"ProjectProtectedActionResultDto", string(data)}, " ")
}
