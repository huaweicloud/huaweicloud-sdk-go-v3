package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectProtectedActionResultApiDto struct {

	// **参数解释：** 事件名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Action *string `json:"action,omitempty"`

	// **参数解释：** 是否启用。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 用户列表。
	Users *[]UserBasicDto `json:"users,omitempty"`

	// **参数解释：** 成员组列表。
	UserTeams *[]UserTeamBasicDto `json:"user_teams,omitempty"`

	// **参数解释：** 角色列表。
	Roles *[]RoleBasicDto `json:"roles,omitempty"`

	// **参数解释：** 操作选择列表。
	AdditionSwitchers *[]ForceActionEnableDto `json:"addition_switchers,omitempty"`
}

func (o ProjectProtectedActionResultApiDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectProtectedActionResultApiDto struct{}"
	}

	return strings.Join([]string{"ProjectProtectedActionResultApiDto", string(data)}, " ")
}
